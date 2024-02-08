package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache	map[string]cacheEntry
	mux		*sync.RWMutex
}

type cacheEntry struct {
	val	[]byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache:	make(map[string]cacheEntry),
		mux: 		&sync.RWMutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock();
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val: val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
  defer c.mux.RUnlock()
	cacheEntry, ok := c.cache[key]

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

	c.mux.Lock()
	defer	c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}
