// Cache that stores retrieved data to prevent
// unnessecary duplicate requests to the API
package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	entries map[string]cacheEntry
	mu *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	value []byte
}

// Create new cache with given time interval
func (ch *Cache) NewCache(interval time.Duration) *Cache {
	newCache := &Cache{}

	// Start reaping goroutine to prune old cache entries
	go newCache.reapLoop(interval)

	return newCache
}

// Add a new entry to the cache
func (ch *Cache) Add(key string, val []byte) {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	// Add new cacheEntry struct
	ch.entries[key] = cacheEntry{
		value: val,
		createdAt: time.Now(),
	}
}

// Get an entry from the cache with find-success flag
func (ch *Cache) Get(key string) ([]byte, bool) {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	entry, found := ch.entries[key]

	return entry.value, found
}

// Remove entries older than interval
func (ch *Cache) reapLoop(interval time.Duration) {
	// Create ticker for time comparison
	ticker := time.NewTicker(interval)

	// Loop iterates every time the interval has passed	
	for range ticker.C {
		// Lock mutex to protect map
		ch.mu.Lock()

		// Delete entries older than interval
		for key, entry := range ch.entries {
			age := time.Since(entry.createdAt)

			if age > interval {
				// Remove old entry
				delete(ch.entries, key)
			}
		}

		// Unlock
		ch.mu.Unlock()
	}
	
}