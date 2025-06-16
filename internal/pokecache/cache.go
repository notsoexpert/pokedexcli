package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data   map[string]CacheEntry
	mu     sync.Mutex
	ticker *time.Ticker
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		data:   make(map[string]CacheEntry),
		ticker: time.NewTicker(interval),
	}
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	elem, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return elem.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		currentTime := <-c.ticker.C
		c.mu.Lock()
		deleteKeys := make([]string, 0)
		for key, val := range c.data {
			lifetime := currentTime.Sub(val.createdAt)
			if lifetime > interval {
				deleteKeys = append(deleteKeys, key)
			}
		}
		for _, key := range deleteKeys {
			delete(c.data, key)
		}
		c.mu.Unlock()
	}
}
