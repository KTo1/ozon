package main

import "time"

// Что выведет программа?
// Программа выведет 6, потому что воркеры запускаются последовательно

func main() {
	timeStart := time.Now()

	ch1 := worker()
	ch2 := worker()

	<-ch1
	<-ch2

	println(int(time.Since(timeStart).Seconds()))
}

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()

	return ch
}
