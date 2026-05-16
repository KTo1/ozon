package main

import "fmt"

func filter(in chan int, f func(int) bool) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			if f(i) {
				out <- i
			}
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

	action := func(val int) bool {
		return val%2 == 0
	}

	for i := range filter(channel, action) {
		fmt.Println(i)
	}
}
