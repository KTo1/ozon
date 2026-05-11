package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go func() {
		defer close(ch1)

		for i := range 10 {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)

		for i := range 10 {
			ch2 <- i
		}
	}()

	start := time.Now()
	res := merge(ctx, ch1, ch2)

	for i := range res {
		fmt.Println(i)
	}

	fmt.Println(time.Since(start))
}

func merge(ctx context.Context, ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ch1:
				if !ok {
					return
				}
				out <- v

				time.Sleep(time.Second)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ch2:
				if !ok {
					return
				}
				out <- v

				time.Sleep(time.Second)
			}
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
