package main

import "time"

type RateLimiter struct {
	Tokens     chan struct{}
	refillTime time.Duration
}

func main() {
}
