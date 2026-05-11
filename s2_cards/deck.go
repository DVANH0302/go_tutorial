package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clutch"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, suite+" of "+val)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Printf("(%d) %s, ", i, card)
	}
	fmt.Println()
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) toByte() []byte {
	return []byte(d.toString())
}

func (d deck) saveTofile(filename string) error {
	dirName := "output"
	err := os.Mkdir(dirName, 0750)

	err = os.WriteFile(filepath.Join(dirName, filename), d.toByte(), 0666)

	return err
}

func newDeckFromFile(filename string) deck {
	dirName := "output"
	data, err := os.ReadFile(filepath.Join(dirName, filename))
	if err != nil {
		log.Fatal(err)
	}
	str_data := string(data)
	str_ls_data := strings.Split(str_data, ", ")
	return deck(str_ls_data)
}

func (d *deck) shuffle() {
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})

}
