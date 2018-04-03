package pool

import (
	"errors"
	"math"
	"time"
)

const MinLifetimeCheckPeriod = 30 * time.Second

type Pool struct {
	config            Config
	items             chan *Item
	hasPending        chan bool
	poolLengthCounter chan int
}

func New(config Config) (*Pool, error) {
	p := &Pool{
		config:            config,
		items:             make(chan *Item, config.Cap),
		hasPending:        make(chan bool),
		poolLengthCounter: make(chan int, 1),
	}

	go func() { p.putPending() }()
	go func() { p.clear() }()

	for i := 0; i < config.InitCap; i++ {
		p.hasPending <- true
	}

	return p, nil
}

func (p *Pool) Get() (*Item, error) {
	var item *Item
	var err error

	timeout := time.After(p.config.Timeout)

	p.hasPending <- true

	select {
	case <-timeout:
		err = errors.New("timeout")
	case item = <-p.items:
		break
	}

	return item, err
}

func (p *Pool) putPending() {
	for {
		<-p.hasPending

		c := <-p.poolLengthCounter

		if len(p.items) == 0 && c < p.config.Cap {
			item, err := p.createItem()
			//todo handle createItem error
			if err == nil {
				p.items <- item
				c++
			}
		}

		p.poolLengthCounter <- c
	}
}

func (p *Pool) createItem() (*Item, error) {
	adapter, err := p.config.Factory.Create()

	if err != nil {
		return nil, err
	}

	item := &Item{
		adapter:      &adapter,
		pool:         p,
		releasedTime: time.Now().UTC(),
	}

	return item, nil
}

func (p *Pool) clear() {

	lifetimeCheckPeriod := p.config.Lifetime / 10
	tick := time.Duration(math.Max(float64(lifetimeCheckPeriod), float64(MinLifetimeCheckPeriod)))

	ticker := time.NewTicker(tick)

	for range ticker.C {
		p.destroyItems()
		p.checkInitCapacity()
	}
}

func (p *Pool) destroyItems() {
	var s []*Item

	for len(p.items) > 0 {
		item := <-p.items
		if item.isReadyForDestroy() {
			p.poolLengthCounter <- <-p.poolLengthCounter - 1
			item.Destroy()
		} else {
			s = append(s, item)
		}
	}

	for _, item := range s {
		p.items <- item
	}
}

func (p *Pool) checkInitCapacity() {
	c := <-p.poolLengthCounter
	p.poolLengthCounter <- c

	delta := p.config.InitCap - c

	for i := 0; i < delta; i++ {
		p.hasPending <- true
	}
}
