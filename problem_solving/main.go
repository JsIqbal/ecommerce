package main

import (
	"container/list"
	"fmt"
)

// LRUCache represents the LRU cache structure.
type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	order    *list.List
}

// CacheItem represents an item stored in the LRU cache.
type CacheItem struct {
	key   string
	value string
}

// NewLRUCache creates a new LRU cache with the specified capacity.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

// Get retrieves the value associated with the given key from the cache.
func (lru *LRUCache) Get(key string) (string, bool) {
	if element, exists := lru.cache[key]; exists {
		// move the accessed item to the front of the list (most recently used).
		lru.order.MoveToFront(element)
		return element.Value.(*CacheItem).value, true
	}

	return "", false
}

// Add adds a new item to the cache or updates the existing one.
func (lru *LRUCache) Add(key, value string) {
	// check if the item already exists.
	if element, exists := lru.cache[key]; exists {
		// update the value and move the item to the front (most recently used).
		element.Value.(*CacheItem).value = value
		lru.order.MoveToFront(element)
	} else {
		// add a new item.
		item := &CacheItem{key, value}
		element := lru.order.PushFront(item)
		lru.cache[key] = element

		// remove the least recently used item if the cache exceeds capacity.
		if len(lru.cache) > lru.capacity {
			lastElement := lru.order.Back()

			if lastElement != nil {
				delete(lru.cache, lastElement.Value.(*CacheItem).key)
				lru.order.Remove(lastElement)
			}
		}
	}
}

func main() {
	// create a new LRU cache with a capacity of 3.
	lruCache := NewLRUCache(3)

	// add some songs to the cache.
	lruCache.Add("song1", "Song One")
	lruCache.Add("song2", "Song Two")
	lruCache.Add("song3", "Song Three")

	// Retrieve songs from the cache.
	if song, exists := lruCache.Get("song2"); exists {
		fmt.Println("Retrieved song from cache:", song)
	} else {
		fmt.Println("Song not found in cache.")
	}

	// add a new song, which should trigger the removal of the least recently used song ("song1").
	lruCache.Add("song4", "Song Four")

	// try to retrieve the removed song ("song1").
	if _, exists := lruCache.Get("song1"); exists {
		fmt.Println("Song 1 found in cache.")
	} else {
		fmt.Println("Song 1 not found in cache (as expected).")
	}
}
