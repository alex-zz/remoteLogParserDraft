package pool

import (
	"time"
)

type Item struct {
	adapter *Adapter
	pool *Pool
	releasedTime time.Time
	idle bool
}

func (i *Item) GetAdapter() *Adapter {
	return i.adapter
}

func (i *Item) Release() {
	i.idle = true;
	i.releasedTime = time.Now().UTC()
	i.pool.c <- i
}

func (i *Item) Destroy() {
	i.idle = false;
	(*i.adapter).Destroy()
}

func (i *Item) isReadyForDestroy() bool {
	expireTime := time.Now().Local().Add(i.pool.lifetime)

	return i.releasedTime.After(expireTime)
}