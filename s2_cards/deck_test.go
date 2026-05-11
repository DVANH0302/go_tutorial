package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("length must be 16, but got %d", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testfile := "_decktesting"
	testfilepath := filepath.Join("output", testfile)
	os.Remove(testfilepath)
	d := newDeck()
	d.saveTofile(testfile)
	loadedDeck := newDeckFromFile(testfile)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards, got %d", len(loadedDeck))
	}
	err := os.Remove(testfilepath)
	if err != nil {
		log.Fatal(err)
	}
}
