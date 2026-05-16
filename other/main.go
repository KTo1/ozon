package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		defer close(ch2)

		ch1 <- 1
		ch1 <- 4
		ch2 <- 2
		ch2 <- 3
	}()

	ch3 := merge(ch1, ch2)

	for v := range ch3 {
		fmt.Println(v)
	}
}

func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}
