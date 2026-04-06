package main

import (
	"fmt"
	"sync"
)

// Задача реализовать кеш с доступом к ключу за О(1)
type LRUCache struct {
	data map[string]string
	mu   sync.RWMutex
	l    int
}

func NewLRUCache(l int) *LRUCache {
	return &LRUCache{
		data: make(map[string]string, l),
		l:    l,
	}
}

func (c *LRUCache) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.data[key]
}

func (c *LRUCache) Put(key, value string) {

}

func main() {
	cache := NewLRUCache(3)
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	key := cache.Get("key3")
	key4 := cache.Get("key4")

	fmt.Println(key)
	fmt.Println(key4)
}
