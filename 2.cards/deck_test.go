package main

import (
	"bytes"
	"os"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	d := createDeck()

	if len(d) != 52 {
		t.Errorf("Expected 52 but got %v", len(d))
	}
	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card is Ace of Hearts but got %v", d[0])
	}
	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected Last card is King of Spades but got %v", d[len(d)-1])
	}

}

func TestDeal(t *testing.T) {
	d := createDeck()
	hand, remainingDeck := deal(d, 3)
	bothHandComparison := bytes.Compare([]byte(hand.toString()), []byte(d[:3].toString()))
	bothRemainingComparison := bytes.Compare([]byte(remainingDeck.toString()), []byte(d[3:].toString()))
	if !(bothHandComparison == 0 && bothRemainingComparison == 0) {
		t.Errorf("Expect Equals but the Dealed deck is not Equal")
	}
}

func TestSaveDeckToFileAndReadDeckFromFile(t *testing.T) {
	d := createDeck()
	fileSaveError := d.saveToDisk("_temp_deck")
	if fileSaveError != nil {
		t.Error("Error : ", fileSaveError)
	}
	loadedDeck := newDeckFromDisk("_temp_deck")
	if len(loadedDeck) != len(d) {
		t.Errorf("Expected lenght of loaded file is 52 but actual length is %v", len(loadedDeck))
	}

	_ = os.Remove("_temp_deck")

}
