// проектирование handler
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Server struct {
	cache *int
	mu    sync.Mutex
	stop  bool
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetLongTask() int {
	s.mu.Lock()
	if s.cache != nil {
		s.mu.Unlock()
		return *s.cache
	}

	s.mu.Unlock()

	res := LongTask()

	s.mu.Lock()
	s.cache = &res
	s.mu.Unlock()

	return res
}

func (s *Server) InvalidateCache() {
	timer := time.NewTicker(5 * time.Second)
	defer timer.Stop()

	for {
		if s.stop {
			break
		}

		select {
		case <-timer.C:
			fmt.Println("invalidating cache")
			res := LongTask()

			s.mu.Lock()
			s.cache = &res
			s.mu.Unlock()
		}
	}
}

func (s *Server) Stop() {
	s.stop = true
}

func LongTask() int {
	//result := callService() // > 5 min
	time.Sleep(10 * time.Second)
	result := rand.Intn(100)

	// process result
	return result
}

func main() {
	// init server

	server := NewServer()
	go server.InvalidateCache()

	handler := func(w http.ResponseWriter, r *http.Request) {
		// write response
		res := server.GetLongTask()
		fmt.Println(res)
		resStr := strconv.Itoa(res)
		w.Write([]byte(resStr))
	}

	http.HandleFunc("/", handler)
	// run server

	http.ListenAndServe(":8888", nil)

	server.Stop()
}
