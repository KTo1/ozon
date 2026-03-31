// Какие проблемы вы видите в этой программе и как их исправить
//  1. Приведение типов
//     непонятно что и когда приводить и можно просто перепутать ключи
//  2. Лоад пусть возвращет что ключт есть
//
// 3. Исправить вызывающий код, сделать проверки на существование ключа
// и тайпассершены
// как вариант сдеалть асершены в методе лоад, но тогда надо каждое значение кастовать ко всем типам

package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func (c *Cache) Store(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

func (c *Cache) Load(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.data[key]
	if !ok {
		return nil, false
	}

	// как пример, но тут придется все типы пробегать)
	nameStr, ok := val.(string)
	if ok {
		return nameStr, true
	}

	return val, ok
}

func main() {
	cache := &Cache{
		data: make(map[string]interface{}),
	}

	cache.Store("Name", "Alice")
	cache.Store("age", 25)

	name, ok := cache.Load("Name")
	if !ok {
		fmt.Println("name not found")
	}

	nameStr, ok := name.(string)
	if !ok {
		fmt.Println("name is not string")
	}
	fmt.Println(nameStr)

	age, ok := cache.Load("age")
	if !ok {
		fmt.Println("age not found")
	}

	ageInt, ok := age.(int)
	fmt.Println(ageInt)

	height, ok := cache.Load("Height")
	if !ok {
		fmt.Println("height not found")
	}
	if height == nil {
		fmt.Println("height is nil")
	}

	width, ok := cache.Load("Width")
	if !ok {
		fmt.Println("width not found")
	}

	widthFloat, ok := width.(float64)
	if !ok {
		fmt.Println("width is not float64")
	}
	fmt.Println(widthFloat)
}
