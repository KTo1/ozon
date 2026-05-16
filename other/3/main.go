package main

import (
	"fmt"
	"sync"
)

func main() {
	numWorkers := 5

	dataCh := make(chan string)
	go func() {
		defer close(dataCh)
		data := getData()
		for i := 0; i < len(data); i++ {
			dataCh <- data[i]
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(dataCh, &wg)
	}

	wg.Wait()
}

func worker(dataCh <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range dataCh {
		if !checkDomain(val) {
			fmt.Printf("%s is bad \n", val)
		}
	}
}

func getData() []string {
	return []string{"1.test", "2.test", "1000.test"}
}

func checkDomain(host string) bool {
	return false
}
