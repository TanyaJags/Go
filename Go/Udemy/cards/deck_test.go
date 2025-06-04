package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newdeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v", deck[0])
	}

	if deck[len(deck)-2] != "Queen of Clubs" {
		t.Errorf("Expected last card to be 'Queen of Clubs', but got %v", deck[len(deck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newdeck()
	err := deck.saveToFile(filename)
	if err != nil {
		t.Errorf("Error saving to file: %v", err)
	}

	loadedDeck := newdeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in the loaded deck, but got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
