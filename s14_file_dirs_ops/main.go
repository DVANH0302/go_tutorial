package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("hello.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file does not exist")

		} else {
			panic(err)
		}
	}

	fmt.Println(info.Name())
	fmt.Println(info.Mode())
}
