// что будет?
// а будет андефайнед бехейвор
// селект выбирает каналы случайным образом,
// оба канала закрыты поэтому запись в ch1 вызовет панику,
// а чтение из ch2 считает зеро вэлью

package main

import "fmt"

func main() {
	ch1 := make(chan string, 1)
	close(ch1)

	ch2 := make(chan struct{})
	close(ch2)

	select {
	case ch1 <- "msg":
		fmt.Println("msg received")
	case <-ch2:
		fmt.Println("channel closed")
	}
}
