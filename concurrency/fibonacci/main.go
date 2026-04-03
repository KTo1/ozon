//## Вопрос
//Реализовать бесконечный генератор Фибоначчи
//0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181

package main

import (
	"fmt"
	"time"
)

func generator() chan int {
	out := make(chan int)

	go func() {
		a, b := 0, 1
		for {
			out <- a
			a, b = b, a+b
			time.Sleep(time.Second)
		}
	}()

	return out
}

func main() {

	ch := generator()

	for {
		i := <-ch
		fmt.Println(i)
	}
}
