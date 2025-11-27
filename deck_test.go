package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 52 {
		t.Errorf("Expected deck length 52, got %d", len(deck))
	}

	if deck[0] != "Ace of Hearts" {
		t.Errorf("Expected first card to be Ace of Hearts, got %s", deck[0])
	}

	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, got %s", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting");
	
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting", true)
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length 52, got %d", len(loadedDeck))
	}
}