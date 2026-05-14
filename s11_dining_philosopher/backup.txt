// package main

// import (
// 	"fmt"
// 	"log"
// 	"math/rand/v2"
// 	"os"
// 	"strconv"
// 	"sync"
// 	"time"
// )

// /*
// 	worker [0,1,2,3,4,5,6]
// 	fork [0,1,2,3,4,5,6]

// 	0 take fork 0, 1
// 	1 take fork 1, 2
// 	2 take fork 2, 3
// 	...
// 	6 take fork 6, 0
// */

// type Worker struct {
// }

// func (w *Worker) pickup_left(id int, fork *sync.Mutex) {
// 	fork.Lock()
// 	// fmt.Printf("Worker %d is picking up LEFT fork\n", id)
// }

// func (w *Worker) pickup_right(id int, fork *sync.Mutex) {
// 	fork.Lock()
// 	// fmt.Printf("Worker %d is picking up RIGHT fork\n", id)
// }

// func (w *Worker) put_down(id int, left_fork *sync.Mutex, right_fork *sync.Mutex) {
// 	left_fork.Unlock()
// 	right_fork.Unlock()
// }

// func (w *Worker) eat(id int, n int, forks []*sync.Mutex, done []int, initial time.Time) {
// 	left_fork := forks[id%n]
// 	right_fork := forks[(id+1)%n]

// 	w.pickup_left(id, left_fork)

// 	w.pickup_right(id, right_fork)

// 	randomTime := time.Duration(rand.IntN(4) + 1)
// 	startTime := time.Since(initial)
// 	fmt.Printf("TIME %.0f - EATING %d: worker %d is eating in %d seconds\n", float64(startTime)/1e9, id, id, randomTime)
// 	time.Sleep(randomTime * time.Second)

// 	w.put_down(id, left_fork, right_fork)

// 	done[id] = 1
// 	endTime := time.Since(initial)
// 	fmt.Printf("TIME %.0f - DONE %d: Worker %d finished from %.0f to %.0f in %d seconds\n", float64(endTime)/1e9, id, id, float64(startTime)/1e9, float64(endTime)/1e9, randomTime)
// }

// func main() {

// 	str_n := os.Args[1]

// 	n, err := strconv.Atoi(str_n)
// 	if err != nil {
// 		log.Fatal("wrong format n", err)
// 	}

// 	fmt.Printf("Initializing %d workers...\n", n)
// 	workers := make([]Worker, n)
// 	done := make([]int, n)
// 	forks := make([]*sync.Mutex, n)
// 	for i := 0; i < n; i++ {
// 		forks[i] = &sync.Mutex{}
// 	}

// 	var wg sync.WaitGroup
// 	t := time.Now()
// 	// the first person eat alone
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		workers[0].eat(0, n, forks, done, t)
// 	}()
// 	wg.Wait()

// 	// do the odd index first
// 	for i := 1; i < n; i += 2 {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			workers[i].eat(i, n, forks, done, t)
// 		}()
// 	}
// 	wg.Wait()

// 	// do the even index last
// 	for i := 2; i < n; i += 2 {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			workers[i].eat(i, n, forks, done, t)
// 		}()
// 	}
// 	wg.Wait()

// 	total := 0
// 	for i := range n {
// 		total += done[i]
// 	}

// 	fmt.Printf("%d People finish eating!\n", total)

// }

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

/*
	worker [0,1,2,3,4,5,6]
	fork [0,1,2,3,4,5,6]

	0 take fork 0, 1
	1 take fork 1, 2
	2 take fork 2, 3
	...
	6 take fork 6, 0
*/

type Worker struct {
}

func (w *Worker) pickup_left(id int, fork *sync.Mutex) {
	fork.Lock()
	fmt.Printf("\n    Worker %d is picking up LEFT fork\n", id)
}

func (w *Worker) pickup_right(id int, fork *sync.Mutex) {
	fork.Lock()
	fmt.Printf("\n    Worker %d is picking up RIGHT fork\n", id)
}

func (w *Worker) put_down(id int, left_fork *sync.Mutex, right_fork *sync.Mutex) {
	left_fork.Unlock()
	right_fork.Unlock()
}

func (w *Worker) eat(id int, n int, forks []*sync.Mutex, done []int, initial time.Time) {
	left_fork := forks[id%n]
	right_fork := forks[(id+1)%n]

	w.pickup_left(id, left_fork)

	randomTime := 3 * time.Second

	startTime := time.Since(initial)
	time.Sleep(randomTime)
	randomTime /= 10e9

	w.pickup_right(id, right_fork)

	fmt.Printf("TIME %.0f - EATING %d: worker %d is eating in %d seconds\n", float64(startTime)/1e9, id, id, randomTime)

	w.put_down(id, left_fork, right_fork)

	done[id] = 1
	endTime := time.Since(initial)
	fmt.Printf("TIME %.0f - DONE %d: Worker %d finished from %.0f to %.0f in %d seconds\n", float64(endTime)/1e9, id, id, float64(startTime)/1e9, float64(endTime)/1e9, randomTime)
}

func (w *Worker) eat_last(id int, n int, forks []*sync.Mutex, done []int, initial time.Time) {
	left_fork := forks[id%n]
	right_fork := forks[(id+1)%n]

	w.pickup_right(id, right_fork)

	randomTime := 3 * time.Second
	startTime := time.Since(initial)

	time.Sleep(randomTime)
	randomTime /= 10e9
	w.pickup_left(id, left_fork)

	fmt.Printf("TIME %.0f - EATING %d: worker %d is eating in %d seconds\n", float64(startTime)/1e9, id, id, randomTime)

	w.put_down(id, left_fork, right_fork)

	done[id] = 1
	endTime := time.Since(initial)
	fmt.Printf("TIME %.0f - DONE %d: Worker %d finished from %.0f to %.0f\n", float64(endTime)/1e9, id, id, float64(startTime)/1e9, float64(endTime)/1e9)
}

func main() {

	str_n := os.Args[1]

	n, err := strconv.Atoi(str_n)
	if err != nil {
		log.Fatal("wrong format n", err)
	}

	fmt.Printf("Initializing %d workers...\n", n)
	workers := make([]Worker, n)
	done := make([]int, n)
	forks := make([]*sync.Mutex, n)
	for i := 0; i < n; i++ {
		forks[i] = &sync.Mutex{}
	}

	var wg sync.WaitGroup
	t := time.Now()
	for i := 0; i < n-1; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			workers[i].eat(i, n, forks, done, t)
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		workers[n-1].eat_last(n-1, n, forks, done, t)
	}()

	wg.Wait()

	total := 0
	for i := range n {
		total += done[i]
	}

	fmt.Printf("%d People finish eating!\n", total)
	fmt.Println(done)
}
