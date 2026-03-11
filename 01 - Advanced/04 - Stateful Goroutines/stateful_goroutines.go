package main

import (
	"fmt"
	"time"
)

type StatefulWorker struct {
	Count int
	Ch    chan int
}

func (sw *StatefulWorker) Start() {
	go func() {
		for {
			select {
			case value := <-sw.Ch:
				sw.Count += value
				fmt.Println("Current count:", sw.Count)
			}
		}
	}()
}

func (sw *StatefulWorker) Send(value int) {
	sw.Ch <- value
}

func main() {
	worker := &StatefulWorker{
		Ch: make(chan int),
	}

	worker.Start()

	for i := range 5 {
		worker.Send(i)
		time.Sleep(500 * time.Millisecond)
	}
}
