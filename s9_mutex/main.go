package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// variable for bank balance
	// print out starting values
	// define weekly revenue
	// loop through 52 weeks and print out how much is made; keep a running total
	// print out final balance
	balance := 0.0
	weeklyRevenue := 100.0
	var mu sync.Mutex
	for i := 0; i < 52; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			balance += weeklyRevenue
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("Final balance:", balance)
}
