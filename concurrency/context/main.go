package main

// У нас есть функция requestThatCostsMoney вызов которой стоит денег (например платное АПИ).
// нужно написать обертку RequestThatCostsMoney так, что бы при параллельных вызовах с одним и
// тем же аргументом, метод requestThatCostsMoney вызвался только один раз.

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mo := NewMoneyOptimization()
	userID := "vasya"

	parralelQuery := 1000

	wg := &sync.WaitGroup{}
	wg.Add(parralelQuery)

	for i := 0; i < parralelQuery; i++ {
		go func() {
			defer wg.Done()
			mo.RequestThatCostsMoney(userID)
		}()

	}

	wg.Wait()
	// тут еще может быть множество параллельно запускаемых запросов
}

func requestThatCostsMoney(userID string) (response string) {
	_ = userID
	// функция, которая стоит денег
	time.Sleep(10 * time.Millisecond)
	return "some"
}

type Response struct {
	response string
	ch       chan struct{}
}

type MoneyOptimization struct {
	cache map[string]*Response
	mu    sync.Mutex
}

func NewMoneyOptimization() *MoneyOptimization {
	return &MoneyOptimization{
		cache: make(map[string]*Response),
	}
}

func (o *MoneyOptimization) RequestThatCostsMoney(userID string) (response string) {
	o.mu.Lock()
	if res, ok := o.cache[userID]; ok {
		o.mu.Unlock()

		<-res.ch

		fmt.Println("Unpaid query")
		return res.response
	}

	o.cache[userID] = &Response{ch: make(chan struct{})}
	o.mu.Unlock()

	res := requestThatCostsMoney(userID)
	fmt.Println("paid query")

	o.mu.Lock()
	o.cache[userID].response = res
	o.mu.Unlock()

	close(o.cache[userID].ch)

	o.mu.Lock()
	delete(o.cache, userID)
	o.mu.Unlock()

	return res
}
