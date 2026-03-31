// Требуется доработать так, что бы:
// код работал в конкурентной среде
// при долгом ожидании отвалимваля по таймауту
// на консоль печаталось время выполенения запроса

package main

//import (
//	"log"
//	"math/rand"
//	"time"
//)
//
//var counter int64
//
//func SimulateRequest() int64 {
//	time.Sleep(time.Duration(rand.Int63n(5)) * time.Second)
//	counter++
//
//	return counter
//}
//
//func main() {
//	val := SimulateRequest()
//	log.Printf("Значение счетчика %d\n", val)
//}
