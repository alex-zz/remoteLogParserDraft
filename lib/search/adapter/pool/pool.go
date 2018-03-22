package pool

import (
	"time"
	"errors"
)

type Pool struct {
	cap int
	lifetime time.Duration
	timeout time.Duration
	factory *Creator
	poolLength int
	c chan *Item
	pendingItem bool
}

func New(cap int, initCap int, lifetime time.Duration, timeout time.Duration, factory Creator) (*Pool, error) {
	p := &Pool{
		cap: cap,
		c: make(chan *Item, cap),
		lifetime: lifetime,
		timeout: timeout,
		factory: &factory,
		pendingItem: false,
	}

	for i := 0; i < initCap; i++ {
		p.put()
	}

	go func() {p.putPending()}()
	go func() {p.clear()}()

	return p, nil
}

func (p *Pool) Get() (*Item, error) {
	var item *Item
	var err error

	timeout := time.After(p.timeout)

	WaitItem:
		for {
			select {
				case <- timeout:
					err = errors.New("timeout")
					break WaitItem

				case item = <- p.c:
					item.idle = false
					break WaitItem

				default:
					p.pendingItem = true
			}
		}

	return item, err
}

func (p *Pool) putPending() {
	for {
		if p.pendingItem {
			err := p.put()

			if err == nil {
				p.pendingItem = false
			}
		}
	}
}

func (p *Pool) put() (error) {

	if p.poolLength < p.cap {
		adapter, err := (*p.factory).Create()

		if err != nil {
			return err
		}

		item := &Item{
			adapter: &adapter,
			pool: p,
			releasedTime: time.Now().UTC(),
			idle: true,
		}

		p.c <- item

		p.poolLength++
	}

	return errors.New("cannot put item")
}

func (p *Pool) clear() {

	//todo set min limit
	ticker := time.NewTicker(p.lifetime / 10)

	for range ticker.C {

		var s []*Item

		for len(p.c) > 0 {
			item := <- p.c
			if item.isReadyForDestroy() {
				item.Destroy()
				p.poolLength--
			} else {
				s = append(s, item)
			}
		}

		for _, item := range s {
			p.c <- item
		}
	}
}