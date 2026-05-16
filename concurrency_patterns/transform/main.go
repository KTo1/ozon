package main

import "fmt"

func transform(in chan int, f func(int) int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			out <- f(i)
		}
	}()

	return out
}

func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)

		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	action := func(val int) int {
		return val - val
	}

	for i := range transform(channel, action) {
		fmt.Println(i)
	}
}
