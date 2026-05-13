package main

import (
	"fmt"
	"sync"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		printSomething("This is the first thing")
		defer wg.Done()
	}()
	go func() {
		printSomething("This is the second thing")
		defer wg.Done()
	}()

	wg.Wait()
	fmt.Println("waiting here")
}
