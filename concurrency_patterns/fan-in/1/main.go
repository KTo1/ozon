// реализовать метод мержа каналов

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func merge(ctx context.Context, channels ...chan int64) <-chan int64 {
	result := make(chan int64)

	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	for _, channel := range channels {
		go func() {
			defer wg.Done()

			for {
				select {
				case v, ok := <-channel:
					if !ok {
						return
					}

					result <- v
				case <-ctx.Done():
					fmt.Println(ctx.Err())
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	channels := make([]chan int64, 1000)
	for i := range 1000 {
		channels[i] = make(chan int64)
	}

	for i := range channels {
		go func(i int) {
			channels[i] <- int64(i)
			close(channels[i])
		}(i)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	start := time.Now()
	for v := range merge(ctx, channels...) {
		println(v)
	}

	fmt.Println(time.Since(start))
}
