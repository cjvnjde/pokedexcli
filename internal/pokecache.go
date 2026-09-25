package internal

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

func NewCache(duration time.Duration) *Cache {
	cache := Cache{
		cache: map[string]cacheEntry{},
		mu:    &sync.Mutex{},
	}
	go cache.reapLoop(duration)
	return &cache
}

func (c *Cache) Add(key string, value []byte) {
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) (data []byte, ok bool) {
	if v, ok := c.cache[key]; ok {
		return v.val, true
	}

	return []byte{}, false
}

func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)

	for {
		<-ticker.C
		c.mu.Lock()
		timeNow := time.Now().Add(-duration)
		for key, value := range c.cache {
			if !(timeNow.Compare(value.createdAt) < 0) {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
