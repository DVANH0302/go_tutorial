package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := make(map[string]string)
	colors = map[string]string{
		"red":   "#ff0000",
		"black": "#4b4234",
	}
	// colors["white"] = "#ffff"

	// delete(colors, "red")

	printMap(colors)

}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
