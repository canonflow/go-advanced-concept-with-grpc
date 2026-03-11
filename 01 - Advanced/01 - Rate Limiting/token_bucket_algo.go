package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	Tokens     chan struct{}
	refillTime time.Duration
}

func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rateLimiter := &RateLimiter{
		Tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}

	// First, add tokens based on the rateLimit params
	for range rateLimit {
		rateLimiter.Tokens <- struct{}{}
	}

	// Then, use goroutine to start the ticker to refill the token if there is an empty slot in bucket
	go rateLimiter.StartRefill()

	return rateLimiter
}

func (rl *RateLimiter) StartRefill() {
	// Create a new token in certain times
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			select {
			case rl.Tokens <- struct{}{}:
			default:
			}
		}
	}
}

func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.Tokens:
		// Consume the token in the bucket if available
		return true
	default:
		return false
	}
}

func main() {
	rateLimiter := NewRateLimiter(5, time.Second)

	for range 10 {
		if rateLimiter.Allow() {
			fmt.Println("Request Allowed")
		} else {
			fmt.Println("Request Denied")
		}

		time.Sleep(200 * time.Millisecond)
	}

	/*
		1 -> 0ms		First request allowed 	(5 tokens left)
		2 -> 200ms		Second request allowed 	(4 tokens left)
		3 -> 400ms		Third request allowed 	(3 tokens left)
		4 -> 600ms		Fourth request allowed	(2 tokens left)
		5 -> 800ms		Fifth request allowed	(1 tokens left)
		6 -> 1000ms		Sixt request allowed	(1 tokens left, the StartRefill executes and adds one more token)
		7 -> 1200ms		Seven request denied	(0 tokens left)
		8 -> 1400ms		Eight request denied	(0 tokens left)
		9 -> 1600ms		Ninth request denied	(0 tokens left)
		10 -> 1800ms	Tenth request denied	(0 tokens left)
	*/
}
