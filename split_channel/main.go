package main

import (
	"fmt"
	"sync"
)

func split(in chan int) (out1 chan int, out2 chan int) {
	out1 = make(chan int)
	out2 = make(chan int)

	wg1 := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}

	wg1.Add(1)
	wg2.Add(1)
	go func() {
		defer wg1.Done()
		defer wg2.Done()

		for i := range in {
			j := i
			wg1.Add(1)
			go func(j int) {
				defer wg1.Done()
				out1 <- j
			}(j)

			wg2.Add(1)
			go func(j int) {
				defer wg2.Done()
				out2 <- j
			}(j)
		}
	}()

	go func() {
		wg1.Wait()
		close(out1)
	}()

	go func() {
		wg2.Wait()
		close(out2)
	}()

	return out1, out2
}

func main() {

	in := make(chan int)

	go func() {
		defer close(in)

		for i := range 5 {
			in <- i
		}
	}()

	out1, out2 := split(in)

	for v := range out1 {
		fmt.Println(v)
	}

	fmt.Println("---------------")
	for v := range out2 {
		fmt.Println(v)
	}
}
