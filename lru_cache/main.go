package main

import (
	"container/list"
	"fmt"
	"sync"
)

// Задача реализовать кеш с доступом к ключу за О(1)
type LRUCache struct {
	data map[string]*list.Element
	cap  int
	list *list.List
	mu   sync.Mutex
}

type entry struct {
	key string
	val string
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		data: make(map[string]*list.Element),
		cap:  cap,
		list: list.New(),
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if v, ok := c.data[key]; ok {
		c.list.MoveToFront(v)
		return v.Value.(*entry).val, true
	}

	return "", false
}

func (c *LRUCache) Put(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if v, ok := c.data[key]; ok {
		c.list.MoveToFront(v)
		v.Value.(*entry).val = value
		return
	}

	if c.list.Len() >= c.cap {
		back := c.list.Back()
		if back != nil {
			delete(c.data, back.Value.(*entry).key)
			c.list.Remove(back)
		}
	}

	e := c.list.PushFront(&entry{key, value})
	c.data[key] = e
}

func main() {
	cache := NewLRUCache(3)
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	key, _ := cache.Get("key3")
	key4, _ := cache.Get("key4")

	fmt.Println(key)
	fmt.Println(key4)
}
