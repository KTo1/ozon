package main

import (
	"fmt"
	"sync"
	"time"
)

type call struct {
	err  error
	val  any
	done chan struct{}
}

type SingleFlight struct {
	mu    sync.Mutex
	calls map[string]*call
}

func NewSingleFlight() *SingleFlight {
	return &SingleFlight{
		calls: make(map[string]*call),
	}
}

func (s *SingleFlight) Do(key string, action func() (any, error)) (any, error) {
	s.mu.Lock()
	if call, found := s.calls[key]; found {
		s.mu.Unlock()
		return s.wait(call)
	}

	call := &call{
		done: make(chan struct{}),
	}

	s.calls[""] = call
	s.mu.Unlock()

	go func() {
		defer func() {
			s.mu.Lock()
			close(call.done)
			delete(s.calls, key)
			s.mu.Unlock()
		}()

		call.val, call.err = action()
	}()

	return s.wait(call)
}

func (s *SingleFlight) wait(call *call) (any, error) {
	<-call.done
	return call.val, call.err
}

func main() {
	const inFlightRequest = 5
	var wg sync.WaitGroup
	wg.Add(inFlightRequest)

	singleFlight := NewSingleFlight()
	const key = "some_key"

	for i := 0; i < inFlightRequest; i++ {
		go func() {
			defer wg.Done()

			val, err := singleFlight.Do(key, func() (any, error) {
				fmt.Println("Single flight")
				time.Sleep(5 * time.Second)
				return "result", nil
			})

			fmt.Println(i, "=", val, err)
		}()
	}

	wg.Wait()
}
