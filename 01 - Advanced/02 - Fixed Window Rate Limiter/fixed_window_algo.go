package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter2 struct {
	Mu        sync.Mutex
	Count     int
	Limit     int
	Window    time.Duration
	ResetTime time.Time
}

func NeweFixedWindowLimiter(limit int, window time.Duration) *RateLimiter2 {
	return &RateLimiter2{
		Limit:  limit,
		Window: window,
	}
}

func (rl *RateLimiter2) Allow() bool {
	rl.Mu.Lock()
	defer rl.Mu.Unlock()

	now := time.Now()

	// Reset the rate limiter
	if now.After(rl.ResetTime) {
		rl.ResetTime = now.Add(rl.Window)
		rl.Count = 0
	}

	if rl.Count < rl.Limit {
		rl.Count++
		return true
	}

	return false
}

func main() {
	rateLimiter := NeweFixedWindowLimiter(3, 1*time.Second)

	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			if rateLimiter.Allow() {
				fmt.Println("Request allowed")
			} else {
				fmt.Println("Request denied")
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
