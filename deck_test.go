package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52, but got %v", len(d))
	}
}

func TestFirstElement(t *testing.T) {
	d := newDeck()
	if d[0] != "Ace Of Hearts" {
		t.Errorf("Expected to be Ace Of Hearts, but received the following instead: %v", d[0])
	}
}
func TestLastElement(t *testing.T) {
	d := newDeck()
	if d[len(d)-1] != "King Of Clubs" {
		t.Errorf("Expected to be King Of Clubs, but received the following instead: %v", d[0])
	}
}

func TestSaveToFileAndRetrieveFromNewDeck(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedD := newDeckFromFile("_decktesting")
	if len(loadedD) != 52 {
		t.Errorf("Length is wrong when loaded from file, length needs to be 52 but is %v", len(d))
	}
	os.Remove("_deckTesting")
}
