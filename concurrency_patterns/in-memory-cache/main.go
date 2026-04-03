// реализовать инмемори кеш
// при инициализации кеша можно передать начальный размер
// порассуждать про размер кеша
// при записи мы лочим весь кеш, если кэш будет большой то надо делать шарды
// 	функуия д.б. устойчива к коллизиям (на два ключа один хэш) и идемпотентна
// если число шардов меняется то
//		балансировка
//		согласованное хеширование
// Можно методы Set Get вынести в шард, красивее будет

package main

import (
	"fmt"
	"hash/fnv"
	"sync"
)

type Cache interface {
	Set(k, v string)
	Get(k string) (string, bool)
}

type Shard struct {
	data map[string]string
	mu   sync.RWMutex
}

type MemoryCache struct {
	shards []*Shard
}

// можно передать начальный размер
func NewMemoryCache(size int) *MemoryCache {
	shards := make([]*Shard, 0, size)
	for range size {
		shards = append(shards, &Shard{
			data: make(map[string]string),
		})
	}

	return &MemoryCache{
		shards: shards,
	}
}

func (c *MemoryCache) Set(k, v string) {
	n := hash(k) % len(c.shards)

	c.shards[n].mu.Lock()
	defer c.shards[n].mu.Unlock()

	c.shards[n].data[k] = v
}

func (c *MemoryCache) Get(k string) (string, bool) {
	n := hash(k) % len(c.shards)

	c.shards[n].mu.RLock()
	defer c.shards[n].mu.RUnlock()

	v, ok := c.shards[n].data[k]

	return v, ok
}

func main() {
	cache := NewMemoryCache(3)
	cache.Set("price", "10000")

	val, ok := cache.Get("price")
	if !ok {
		fmt.Println("cache miss")
	}

	println(val)
}

func hash(key string) int {
	h := fnv.New32()
	h.Write([]byte(key))

	return int(h.Sum32())
}
