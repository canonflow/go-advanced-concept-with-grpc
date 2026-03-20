package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// kill -s SIGINT PID
	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sign := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Notify the channel on interrupt or terminate signals
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGUSR1)

	go func() {
		sig := <-sign
		fmt.Println("Received signal:", sig)
		done <- true
	}()

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Stopping work due to signal.")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(time.Second)
			}
		}

		// for sig := range sign {
		// 	switch sig {
		// 	case syscall.SIGINT:
		// 		fmt.Println("Received SIGINT (Interrupt)")
		// 	case syscall.SIGTERM:
		// 		fmt.Println("Received SIGTERM (Terminate)")
		// 	case syscall.SIGHUP:
		// 		fmt.Println("Received SIGHUP (Hangup)")
		// 	case syscall.SIGUSR1:
		// 		fmt.Println("Received SIGUSR1 (User defined Signal 1)")
		// 		fmt.Println("User defined function is executed")
		// 		continue
		// 	}
		// 	fmt.Println("Graceful exit.")
		// 	os.Exit(0)
		// }
	}()

	// Simulate some work
	// fmt.Println("Working ...")
	for {
		time.Sleep(time.Second)
	}
}
