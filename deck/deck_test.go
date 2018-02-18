package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected Deck length of 16 got, %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades got, %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card Four of Clubs got, %v", len(d)-1)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)
	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)
	if len(loadedDeck) != 16 {
		t.Errorf("Expecting 16 card in Desk, got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
