// Реализация ЛФУ кеша на массивах
package main

import "fmt"

type entryLFU struct {
	key   int
	value int
	freq  int
}

type LFUCache struct {
	cap   int
	items []entryLFU
}

func NewLFUCache(cap int) *LFUCache {
	return &LFUCache{
		cap:   cap,
		items: []entryLFU{},
	}
}

func (c *LFUCache) Get(key int) (int, bool) {
	for i, e := range c.items {
		if e.key == key {
			c.items[i].freq++
			return e.value, true
		}
	}

	return 0, false
}

func (c *LFUCache) Put(key int, value int) {
	// обновляем существующий
	for i, e := range c.items {
		if e.key == key {
			c.items[i].value = value
			c.items[i].freq++
			return
		}
	}

	// новый элемент
	if len(c.items) < c.cap {
		c.items = append(c.items, entryLFU{key, value, 1})
		return
	}

	// место кончилось – ищем наименьшую частоту
	minIdx := 0
	for i := 1; i < len(c.items); i++ {
		if c.items[i].freq < c.items[minIdx].freq {
			minIdx = i
		}
	}

	// заменяем
	c.items[minIdx] = entryLFU{key, value, 1}
}

func main() {
	lfu := NewLFUCache(3)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Put(3, 3)
	lfu.Get(1)              // freq(1) = 2
	lfu.Get(1)              // freq(1) = 3
	lfu.Put(4, 4)           // вытеснится элемент с freq=1 (например, 2 или 3)
	fmt.Println(lfu.Get(2)) // nil, false
}
