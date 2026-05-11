package main

import "fmt"

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type englishBot struct {
}

type spanishBot struct {
}

func (eb englishBot) getGreeting() string {
	return "Hi there"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
