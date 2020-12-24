package generic

import (
	"context"
	"sync"
	"time"
)

type simpleCache struct {
	items  sync.Map
	cancel func()
}

type item struct {
	data    interface{}
	expires int64
}

func NewCache(garbageCollect time.Duration) Cache {
	cache := &simpleCache{
		items: sync.Map{},
	}
	if garbageCollect > 0 {
		go func() {
			ticker := time.NewTicker(garbageCollect)
			ctx, cancel := context.WithCancel(context.Background())
			cache.cancel = cancel
			defer cancel()
			for {
				select {
				case <-ctx.Done():
					ticker.Stop()
					return
				case <-ticker.C:
					now := time.Now().UnixNano()
					cache.items.Range(func(key, value interface{}) bool {
						item := value.(item)
						if item.expires > 0 && now > item.expires {
							cache.items.Delete(key)
						}
						return true
					})
				}
			}
		}()
	}
	return cache
}

func (c *simpleCache) Get(key interface{}) (interface{}, bool) {
	obj, exists := c.items.Load(key)

	if !exists {
		return nil, false
	}

	item := obj.(item)

	if item.expires > 0 && time.Now().UnixNano() > item.expires {
		return nil, false
	}

	return item.data, true
}

func (c *simpleCache) Exists(key interface{}) bool {
	_, ok := c.Get(key)
	return ok
}

func (c *simpleCache) Set(key interface{}, value interface{}, duration time.Duration) {
	var expires int64

	if duration > 0 {
		expires = time.Now().Add(duration).UnixNano()
	}
	c.items.Store(key, item{
		data:    value,
		expires: expires,
	})
}

func (c *simpleCache) Range(f func(key, value interface{}) bool) {
	now := time.Now().UnixNano()

	fn := func(key, value interface{}) bool {
		item := value.(item)

		if item.expires > 0 && now > item.expires {
			return true
		}

		return f(key, item.data)
	}

	c.items.Range(fn)
}

func (c *simpleCache) Delete(key interface{}) {
	c.items.Delete(key)
}

func (c *simpleCache) Len() int {
	count := 0
	c.Range(func(key, value interface{}) bool {
		if value != nil {
			count++
		}
		return true
	})
	return count
}

func (c *simpleCache) Close() {
	c.cancel()
}
