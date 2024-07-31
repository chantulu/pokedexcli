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
    entries map[string]cacheEntry
    mutex   sync.Mutex
}

// NewCache creates a new Cache with a specified reap interval
func NewCache(interval time.Duration) *Cache {
    cache := &Cache{
        entries: make(map[string]cacheEntry),
    }
    go cache.reapLoop(interval)
    return cache
}

// Add adds a new entry to the cache
func (c *Cache) Add(key string, val []byte) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    c.entries[key] = cacheEntry{
        createdAt: time.Now(),
        val:       val,
    }
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mutex.Lock()
		now := time.Now()
		for key,val := range c.entries{
			diff := now.Sub(val.createdAt)
			if diff >= interval {
				delete(c.entries, key)
			}
		}
		c.mutex.Unlock()
	}
}