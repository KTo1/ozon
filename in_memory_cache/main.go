package main

import (
	"fmt"
	"hash/fnv"
	"sync"
)

type ICache interface {
	Set(key string, value string)
	Get(key string) (string, bool)
}

var _ ICache = (*MemoryCache)(nil)

type MemoryCache struct {
	Shards []Shard
}

type Shard struct {
	mu   sync.Mutex
	data map[string]string
}

func NewMemoryCache(num int) *MemoryCache {
	shards := make([]Shard, 0, num)

	for i := 0; i < num; i++ {
		shards = append(shards, Shard{data: make(map[string]string)})
	}

	return &MemoryCache{
		Shards: shards,
	}
}

func (m *MemoryCache) Set(key string, value string) {
	shardID := hash(key) % len(m.Shards)
	m.Shards[shardID].Set(key, value)
}

func (m *MemoryCache) Get(key string) (string, bool) {
	shardID := hash(key) % len(m.Shards)
	return m.Shards[shardID].Get(key)
}

func (s *Shard) Set(key string, value string) {
	defer s.mu.Unlock()

	s.mu.Lock()
	s.data[key] = value
}

func (s *Shard) Get(key string) (string, bool) {
	defer s.mu.Unlock()
	s.mu.Lock()

	v, ok := s.data[key]

	return v, ok
}

func main() {
	cache := NewMemoryCache(5)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key4", "value4")
	cache.Set("key5", "value5")

	v, ok := cache.Get("key1")
	fmt.Println("key 1: ", v, ok)

	v, ok = cache.Get("key2")
	fmt.Println("key 1: ", v, ok)

	v, ok = cache.Get("key3")
	fmt.Println("key 1: ", v, ok)
}

// private

func hash(key string) int {
	h := fnv.New32()
	h.Write([]byte(key))

	return int(h.Sum32())
}
