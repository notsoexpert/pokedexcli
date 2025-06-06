package pokecache

import {
	"fmt"
	"time"
	"sync"
}

type Cache struct {
	data 		map[string]cacheEntry
	mu			sync.Mutex
	ticker		time.Ticker
}

type CacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(interval time.Duration) Cache {
	var cache Cache{
		data:	  	make(map[string]cacheEntry)
		ticker:		time.NewTicker(interval)
	}
	go cache.reapLoop(interval)
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = Cache{
		createdAt: 	time.Now()
		val: 		val
	}
	return cache
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	elem, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return elem, true
}

func (c Cache) reapLoop(interval time.Duration) {
	
}