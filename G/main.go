package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	load := make(chan int)
	go func() {
		for i := 1; i < 11; i++ {
			load <- i
		}

		close(load)
	}()

	wg := sync.WaitGroup{}
	wc := runtime.NumCPU()
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go func(gn int) {
			defer wg.Done()
			for i := range load {
				fmt.Println("gn", gn, i*i)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("done")
}
