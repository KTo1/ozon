// Какие проблемы вы видите в этой программе и как их исправить
package main

//import (
//	"fmt"
//	"sync"
//)
//
//type Cache struct {
//	data map[string]interface{}
//	mu   sync.RWMutex
//}
//
//func (c *Cache) Store(key string, value interface{}) {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//
//	c.data[key] = value
//}
//
//func (c *Cache) Load(key string) interface{} {
//	c.mu.RLock()
//	defer c.mu.RUnlock()
//
//	return c.data[key]
//}
//
//func main() {
//	cache := &Cache{
//		data: make(map[string]interface{}),
//	}
//
//	cache.Store("Name", "Alice")
//	cache.Store("age", 25)
//
//	name := cache.Load("Name").(string)
//	fmt.Println(name)
//
//	age := cache.Load("age").(int)
//	fmt.Println(age)
//
//	height := cache.Load("Height")
//	if height == nil {
//		fmt.Println("height is nil")
//	}
//
//	width := cache.Load("Width").(float64)
//	fmt.Println(width)
//}
