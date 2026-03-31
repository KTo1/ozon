package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{data: make(map[string]string)}
}

func (c *ConcurrentMap) GetOrCreate(key string, value string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if v, ok := c.data[key]; ok {
		return v
	}

	c.data[key] = value

	return value
}

func (c *ConcurrentMap) GetOrCreateFast(key string, value string) string {
	c.mu.RLock()
	v, ok := c.data[key]
	c.mu.RUnlock()

	if ok {
		return v
	}

	return c.GetOrCreate(key, value)
}

func main() {
	cm := NewConcurrentMap()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		val := cm.GetOrCreate("key1", "value1")
		fmt.Println("goroutine 1: ", val)
	}()

	go func() {
		defer wg.Done()

		val := cm.GetOrCreate("key1", "value2")
		fmt.Println("goroutine 2: ", val)
	}()

	wg.Wait()

}
