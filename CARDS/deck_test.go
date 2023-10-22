package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	elemnt_1 := "Ace of Spades"

	if d[0] != elemnt_1 {
		t.Errorf("expected first element to be [ %v ], but got -> %v", elemnt_1, d[0])
	}

	last_element := "Four of Clubs"

	if d[len(d)-1] != last_element {
		t.Errorf("Expected last element to be [ %v ], but got -> %v", last_element, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromDisk(t *testing.T) {
	testing_file := "_decktesting"

	os.Remove(testing_file)

	d := newDeck()
	d.saveToFile(testing_file)

	loadedd := newDeckFromFile(testing_file)

	deckLen := 16
	if len(loadedd) != deckLen {
		t.Errorf("Expected length of %d, but got -> %v", deckLen, len(loadedd))
	}

	os.Remove(testing_file)
}
