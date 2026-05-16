package main

import "time"

type RateLimiter struct {
	closeCh     chan struct{}
	closeDoneCh chan struct{}
	leakyBucket chan struct{}
}

func NewRateLimiter(limit int, period time.Duration) *RateLimiter {
	limiter := &RateLimiter{
		closeCh:     make(chan struct{}),
		closeDoneCh: make(chan struct{}),
		leakyBucket: make(chan struct{}, limit),
	}

	leakInterval := period.Nanoseconds() / int64(limit)
	go limiter.startperiodleak(time.Duration(leakInterval))

	return limiter
}

func (r *RateLimiter) startperiodleak(interval time.Duration) {
	timer := time.NewTimer(interval)
	defer func() {
		timer.Stop()
		close(r.closeDoneCh)
	}()

	for {
		select {
		case <-r.closeCh:
			return
		default:
		}

		select {
		case <-r.closeCh:
			return
		case <-timer.C:
			<-r.leakyBucket
		default:
		}
	}
}

func (r *RateLimiter) Allow() bool {
	select {
	case r.leakyBucket <- struct{}{}:
		return true
	default:
		return false
	}
}

func (r *RateLimiter) Shutdown() {
	close(r.closeCh)
	<-r.closeDoneCh
}

func main() {

}
