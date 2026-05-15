package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Barber struct {
	seatChan chan *Customer
}

type Customer struct {
	id       int
	seatChan chan *Customer
}

const SLOTS = 5

func main() {
	seatChan := make(chan *Customer, SLOTS) // 5 slots

	barber := Barber{
		seatChan: seatChan,
	}

	customerCount := 0
	timeoutChan := time.After(10 * time.Second)

	var wg sync.WaitGroup

	// barber open the shop
	wg.Add(1)
	go func() {
		defer wg.Done()
		barber.cut()
	}()

	for {
		gap := rand.IntN(600)
		time.Sleep(time.Duration(gap) * time.Millisecond)

		select {

		case <-timeoutChan:
			fmt.Printf("Time is END! Barber will try finish cutting for %d people waiting.\n", len(seatChan))
			close(seatChan)
			wg.Wait()
			fmt.Println("Finished")
			return

		case <-time.After(time.Duration(gap)):

			newCustomer := Customer{id: customerCount, seatChan: seatChan}
			customerCount++

			newCustomer.trySeating()
		}

	}
}

func (c *Customer) trySeating() {
	select {
	case c.seatChan <- c:
		color.Yellow("Customer %d seat!", c.id)
	default:
		color.Red("Customer %d go away because the shop is busy!", c.id)
	}
}

func (b *Barber) cut() {

	for c := range b.seatChan {
		gap := rand.IntN(600) + 1000
		color.Green("Barber is cutting for CUSTOMER %d in %dms", c.id, gap)
		time.Sleep(time.Duration(gap) * time.Millisecond)
	}

}
