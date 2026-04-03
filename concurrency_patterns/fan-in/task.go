// реализовать метод мержа каналов

package main

//func main() {
//	channels := make([]chan int64, 10)
//	for i := range 10 {
//		channels[i] = make(chan int64)
//	}
//
//	for i := range channels {
//		go func(i int) {
//			channels[i] <- int64(i)
//			close(channels[i])
//		}(i)
//	}
//
//	for v := range merge(channels...) {
//		println(v)
//	}
//}
