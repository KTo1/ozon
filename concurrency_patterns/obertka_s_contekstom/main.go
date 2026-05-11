package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const DefaultTimeOut = 5 * time.Second

func randomWait() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Second)
}

func notRandomWait(ctx context.Context) {
	ch := make(chan int)

	go func() {
		defer close(ch)
		randomWait()
		ch <- 1
	}()

	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, DefaultTimeOut)
		defer cancel()
	}

	select {
	case <-ctx.Done():
		return
	case <-ch:
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	start := time.Now()
	notRandomWait(ctx)
	fmt.Println(time.Since(start))
}
