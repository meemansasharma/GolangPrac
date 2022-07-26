package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected length 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Ace of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loaddeck := newDeckFromFile("_decktesting")
	if len(loaddeck) != 20 {
		t.Errorf("Error Expected length 16 but got %v", len(loaddeck))
	}

	os.Remove("_decktesting")
}
