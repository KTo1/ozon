// что выведет данный код
// 		Ничего, будет дедлок, тикер сработает первый и в строке 30 программа заблокитуеся
// 		так же плохо что тикер не сотановиться его надо завершить
// 		решить проблему можно если добавить буфер 1 каналу
//
//		если нужно что бы горутина доработала, то надо вейтгруппу
//		так же канал по хорошему надо закрыть, но тут неопределенное поведение, т.к. в
//		канал пишут в дрвух местах, вопросы  к коду) зачем такой код нужен

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("release")
		ch <- false
	}()

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case value := <-ch:
			fmt.Println("received", value)
			ticker.Stop()
			return
		case <-ticker.C:
			fmt.Println("tick")
			ch <- true
		}
	}
}
