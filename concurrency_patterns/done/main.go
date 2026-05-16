package main

import (
	"fmt"
	"time"
)

type Worker struct {
	closeCh     chan struct{}
	closeDoneCh chan struct{}
}

func NewWorker() *Worker {
	worker := &Worker{
		closeCh:     make(chan struct{}),
		closeDoneCh: make(chan struct{}),
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		defer func() {
			ticker.Stop()
			close(worker.closeDoneCh)
		}()

		for {
			select {
			case <-worker.closeCh:
				return
			default:
			}

			select {
			case <-worker.closeCh:
			case <-ticker.C:
				fmt.Println("worker")
			}
		}
	}()

	return worker
}

func (w *Worker) Shutdown() {
	//ожидаем когда воркер остановиться и гарантирует что все завершилось
	close(w.closeCh)
	<-w.closeDoneCh
}

func main() {
	worker := NewWorker()
	time.Sleep(5 * time.Second)
	worker.Shutdown()

	fmt.Println("shutdown")
	time.Sleep(5 * time.Second)
}
