// Провести ревью кода
package main

//import (
//	"fmt"
//	"math/rand"
//	"sync"
//	"time"
//)
//
//func LongCalculation(n int) int {
//	secondsToSleep := rand.Float64() * float64(n)
//	time.Sleep(time.Duration(secondsToSleep))
//
//	return n + 1
//}
//
//var cache = map[int]int{}
//
//func CachedLongCalculation(n int) int {
//	var mu sync.Mutex
//
//	mu.Lock()
//	found, ok := cache[n]
//	mu.Unlock()
//
//	if !ok {
//		value := LongCalculation(n)
//
//		mu.Lock()
//		cache[n] = value
//		mu.Unlock()
//
//		return value
//	}
//
//	return found
//}
//
//func main() {
//	nums := []int{5, 10, 22}
//
//	for _, n := range nums {
//		val := CachedLongCalculation(n)
//		fmt.Printf("LongCalculation(%d) = %d\n", n, val)
//	}
//}
