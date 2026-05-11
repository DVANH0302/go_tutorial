package main

import (
	"log"
)

func main() {

	cards := newDeck()

	hand, cards := deal(cards, 5)

	// hand.print()
	// fmt.Println("\n Deck: \n")
	// cards.print()

	e1 := hand.saveTofile("hand.txt")
	if e1 != nil {
		log.Fatal("wrong with saving hand: ", e1)
	}
	e2 := cards.saveTofile("cards.txt")
	if e2 != nil {
		log.Fatal("wrong with saving cards: ", e2)
	}

	new_deck := newDeckFromFile("hand.txt")
	new_deck.print()
	new_deck.shuffle()
	new_deck.print()
}
