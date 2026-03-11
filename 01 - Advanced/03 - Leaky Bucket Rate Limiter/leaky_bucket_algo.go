package main

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	Capacity int
	LeakRate time.Duration
	Tokens   int
	LastLeak time.Time
	Mu       sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		Capacity: capacity,
		LeakRate: leakRate,
		Tokens:   capacity,
		LastLeak: time.Now(),
	}
}

func (lr *LeakyBucket) Allow() bool {
	lr.Mu.Lock()
	defer lr.Mu.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(lr.LastLeak)
	tokensToAdd := int(elapsedTime / lr.LeakRate)
	lr.Tokens += tokensToAdd

	if lr.Tokens > lr.Capacity {
		lr.Tokens = lr.Capacity
	}

	lr.LastLeak = lr.LastLeak.Add(time.Duration(tokensToAdd) * lr.LeakRate)

	fmt.Printf("Tokens added %d, Tokens subtracted %d, Total tokens: %d\n", tokensToAdd, 1, lr.Tokens)
	fmt.Printf("Last leak time: %v\n", lr.LastLeak)
	if lr.Tokens > 0 {
		lr.Tokens--
		return true
	}

	return false
}

func main() {
	leakyBucket := NewLeakyBucket(5, 500*time.Millisecond)

	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if leakyBucket.Allow() {
				fmt.Println("Current time:", time.Now())
				fmt.Println("Request allowed\n")
			} else {
				fmt.Println("Current time:", time.Now())
				fmt.Println("Request denied\n")
			}
			time.Sleep(200 * time.Millisecond)
		}()
	}

	wg.Wait()
}
