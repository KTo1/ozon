// Провести ревью кода
// 1. Лишний анлок мьютекса
// 2. Использовать рвмьютекс
// 3. если код конкурентный то каждый вызов будет создавать свой мьютекс и работать не будет, надо
//	сделать структуру кеша и мьютекс моложить туда,
// 4. сделать конструктор и там инициазировать мапу
// 5. нет стратегии вытеснения кеша, кеш бесконечный
// 6. Для LongCalculation возможно стоит сдеать таймуат
// 7.
// 8. Хит мисс метрики, что бы оценить эффективность кеша
// 9. Сам мейн переделать на горутины

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func LongCalculation(n int) int {
	secondsToSleep := rand.Float64() * float64(n)
	time.Sleep(time.Duration(secondsToSleep))

	return n + 1
}

type Cache struct {
	data map[int]int
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[int]int),
	}
}

var cache = NewCache()

func CachedLongCalculation(n int) int {
	cache.mu.RLock()
	found, ok := cache.data[n]
	cache.mu.RUnlock()

	if !ok {
		value := LongCalculation(n)

		cache.mu.Lock()
		// 7. поскольку эта часть кода не под мьтексом, то может стоит тут еще раз убедиться, что в кеше нет значения
		cache.data[n] = value
		cache.mu.Unlock()

		return value
	}

	//mu.Unlock()

	return found
}

func main() {
	nums := []int{5, 10, 22}

	for _, n := range nums {
		val := CachedLongCalculation(n)
		fmt.Printf("LongCalculation(%d) = %d\n", n, val)
	}
}
