package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type CacheAPI interface {
	Get(key string) ([]byte, bool)
	Add(key string, data []byte)
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheRecords map[string]cacheEntry
	mu           sync.RWMutex
	interval     time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		interval:     interval,
		cacheRecords: make(map[string]cacheEntry),
		mu:           sync.RWMutex{},
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheRecords[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	d, ok := c.cacheRecords[key]
	if !ok {
		return nil, false
	}
	return d.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for k, v := range c.cacheRecords {
			if now.Sub(v.createdAt) > c.interval {
				delete(c.cacheRecords, k)
				fmt.Printf("removed expired record with key: %s\n", k)
			}
		}
		c.mu.Unlock()
	}
}
