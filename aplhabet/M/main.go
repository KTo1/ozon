package main

import (
	"fmt"
	"time"
)

func main() {
	timenow := time.Now()
	_, _ = <-worker(), <-worker()
	fmt.Println(time.Since(timenow).Seconds(), "seconds")
}

func worker() <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		ch <- 1
	}()
	return ch
}
