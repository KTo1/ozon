// Требуется доработать так, что бы:
// код работал в конкурентной среде
//		добавить атомик к каунтер
// при долгом ожидании отвалимваля по таймауту
// 		Добавить контекст, селект и оишбку, а сами вычисления обернуть в горутину и добавить канал, который будет записываться когда это все отработает
// на консоль печаталось время выполенения запроса
//		добавить в метод замер вемени

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var counter int64

func SimulateRequest(ctx context.Context) (int64, error) {
	start := time.Now()
	defer func() {
		fmt.Println("Время выполнения запроса", time.Since(start))
	}()

	ch := make(chan int64)

	go func() {
		time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
		atomic.AddInt64(&counter, 1)
		ch <- counter
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case val, ok := <-ch:
		if !ok {
			return 0, errors.New("канал закрыт")
		}

		return val, nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			_, err := SimulateRequest(ctx)
			if err != nil {
				fmt.Println(err.Error())
			}
		}()
	}

	time.Sleep(time.Second * 15)
	wg.Wait()

	log.Printf("Значение счетчика %d\n", counter)
}
