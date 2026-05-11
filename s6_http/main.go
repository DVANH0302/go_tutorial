package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal("Error", err)
	}
	p := make([]byte, 100000)
	_, err = resp.Body.Read(p)
	if err != nil {
		log.Fatal("Error", err)
	}
	fmt.Println(string(p))
}
