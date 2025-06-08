package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	stop     chan bool
	interval time.Duration
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			c.mu.Lock()
			for key, value := range c.entries {
				if t.Sub(value.createdAt) > c.interval {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		case <-c.stop:
			return
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{entries: make(map[string]cacheEntry), stop: make(chan bool), interval: interval}
	go cache.reapLoop()
	return &cache
}
