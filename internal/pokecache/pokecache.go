package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(interval, c.mux)
	return c
}

func (c *Cache) Add(key string, val []byte, mux *sync.Mutex) {
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string, mux *sync.Mutex) ([]byte, bool) {
	mux.Lock()
	defer mux.Unlock()
	value, ok := c.cache[key]
	return value.val, ok
}

func (c *Cache) reapLoop(interval time.Duration, mux *sync.Mutex) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval, mux)
	}
}

func (c *Cache) reap(interval time.Duration, mux *sync.Mutex) {
	mux.Lock()
	defer mux.Unlock()
	t := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(t) {
			delete(c.cache, k)
		}
	}
}
