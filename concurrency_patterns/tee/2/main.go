package main

import (
	"fmt"
	"sync"
)

func splitChanel(in chan int, n int) []chan int {
	channels := make([]chan int, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan int)
	}

	go func() {
		for i := range in {
			for j := 0; j < n; j++ {
				channels[j] <- i //can be non-blocking
			}
		}

		for i := 0; i < n; i++ {
			close(channels[i])
		}
	}()

	return channels
}

func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	channels := splitChanel(channel, 2)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := range channels[0] {
			fmt.Println(i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := range channels[1] {
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
