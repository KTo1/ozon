// что выведет данный код
package main

//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	ch := make(chan bool)
//
//	go func() {
//		time.Sleep(3 * time.Second)
//		fmt.Println("release")
//		ch <- false
//	}()
//
//	ticker := time.NewTicker(1 * time.Second)
//	for {
//		select {
//		case value := <-ch:
//			fmt.Println("received", value)
//			return
//		case <-ticker.C:
//			fmt.Println("tick")
//			ch <- true
//		}
//	}
//}
