package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, os.Interrupt)

	for range signalChan {
		fmt.Println("\nGoodbye")

		return
	}

}
