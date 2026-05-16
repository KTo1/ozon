package main

import (
	"fmt"
	"time"
)

func OrDone(in <-chan string, done <-chan struct{}) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for {
			select {
			case <-done:
				return
			default:
			}

			select {
			case val, ok := <-in:
				if !ok {
					return
				}

				out <- val
			case <-done:
				return
			}
		}
	}()

	return out
}

func main() {
	channel := make(chan string)

	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			channel <- "Hello World"
		}
	}()

	done := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	for v := range OrDone(channel, done) {
		fmt.Println(v)
	}

	/* заменяем вот эту  бубуйню
	for {
		select {
		case val, ok := <-inputChan:
			if !ok {
				return
			}

			//processing
		case <-done:
			return
		}
	}
	*/
}
