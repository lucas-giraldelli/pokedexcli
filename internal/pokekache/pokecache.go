package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val	[]byte
	createdAt time.Time
}

var mux = &sync.RWMutex{}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	mux.Lock()
	c.cache[key] = cacheEntry{
		val: val,
		createdAt: time.Now().UTC(),
	}
	mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	mux.RLock()
	cacheEntry, ok := c.cache[key]
	mux.RUnlock()

	return cacheEntry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)

	mux.Lock()
	for k, v := range c.cache {
		if v.createdAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
	mux.Unlock()
}
