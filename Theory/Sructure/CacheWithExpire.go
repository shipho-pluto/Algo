package Struct

import (
	"errors"
	"sync"
	"time"
)

type (
	Cache struct {
		db                map[int]Item
		mtx               *sync.RWMutex
		defaultExpiration time.Duration
		cleanupInterval   time.Duration
	}

	Item struct {
		Value      any
		CreatedAt  time.Time
		Expiration int64
	}
)

func NewCache(defaultExpiration time.Duration, cleanupInterval time.Duration) *Cache {
	cache := Cache{
		db:                make(map[int]Item),
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
		mtx:               new(sync.RWMutex),
	}

	if cleanupInterval > 0 {
		cache.StartGC()
	}
	return &cache
}

func (c *Cache) Set(key int, value any, duration time.Duration) {

	var expiration int64

	if duration == 0 {
		duration = c.defaultExpiration
	}

	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.db[key] = Item{
		Value:      value,
		CreatedAt:  time.Now(),
		Expiration: expiration,
	}
}

func (c *Cache) Get(key int) (any, bool) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	item, found := c.db[key]

	if !found {
		return nil, false
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}
	}

	return item.Value, true
}

func (c *Cache) Delete(key int) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if _, found := c.db[key]; !found {
		return errors.New("key not found")
	}
	delete(c.db, key)
	return nil
}

func (c *Cache) StartGC() {
	go c.GC()
}

func (c *Cache) GC() {
	for {
		<-time.After(c.cleanupInterval)
		if c.db == nil {
			return
		}

		if keys := c.ExpiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
		}
	}
}

func (c *Cache) ExpiredKeys() (keys []int) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	for k, i := range c.db {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

func (c *Cache) clearItems(keys []int) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	for _, k := range keys {
		delete(c.db, k)
	}
}
