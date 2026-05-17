package main

import (
	"fmt"
	"time"
)

// Write a program that prints "tick" every 2 seconds and automatically stops after 10 seconds.

func main() {
	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
		case <-done:
			ticker.Stop()
			return
		}
	}
}
