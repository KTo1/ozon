package main

import (
	"fmt"
	"sync"
	"time"
)

type Interfacer interface {
	Get(k, v string) string
}

type Inter struct {
}

func (i *Inter) Get(k, v string) string {
	return k
}

func NewInter(i Interfacer) {

}

func main() {
	counter := make([]int, 0, 1000)

	var wg = sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			counter = append(counter, i)
			time.Sleep(10 * time.Millisecond)
		}()
	}

	wg.Wait()
	fmt.Println(len(counter))
}

func modify(s []int) {
	for i, n := range s {
		s[i] = n * 2
		if i%2 == 0 {
			s = append(s, i*2) // 0, len 4, 6   2 len 5, cap 6
		}
	}

	fmt.Println("in func:", s) // 2, 4, 6, 0, 4
}
