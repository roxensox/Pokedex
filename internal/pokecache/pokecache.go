package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.RWMutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	// Creates a cache instance
	c := Cache{
		// Initializes cache map of string: cacheEntry
		cache: make(map[string]cacheEntry),
		// Creates a pointer to a mutex to control access to the map
		mux: &sync.RWMutex{},
	}
	// Starts reapLoop
	go c.reapLoop(interval)
	// Returns the cache
	return c
}

func (c *Cache) Add(key string, val []byte) {
	// Adds a cacheEntry to the cache
	// Locks the map
	c.mux.Lock()
	// Defers unlock so that it unlocks even on error
	defer c.mux.Unlock()
	// Creates the cacheEntry at key
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// Gets a cacheEntry from the cache map
	// Locks the map
	c.mux.Lock()
	// Defers unlock so that it unlocks even on error
	defer c.mux.Unlock()
	// Attempts to access the data at key
	value, ok := c.cache[key]
	// Returns the value and status of attempt
	return value.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	// Calls the reap method at a set interval
	// Creates a ticker to track the interval
	ticker := time.NewTicker(interval)

	// After the interval, calls reap
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	// Removes cacheEntries that are older than the interval
	// Gets the time that entries should be removed at (questionable phrasing)
	t := time.Now().UTC().Add(-interval)

	// Locks the map
	c.mux.Lock()
	// Defers unlock so that it unlocks even in case of error
	defer c.mux.Unlock()
	// Loops through each key value pair in the map
	for k, v := range c.cache {
		// If the entry was created before the set time, it's deleted
		if v.createdAt.Before(t) {
			delete(c.cache, k)
		}
	}
}
