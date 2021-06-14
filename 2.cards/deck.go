package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Something like class deck extends string slice
type deck []string

func createDeck() deck {
	cards := deck{}
	cardSuites := [4]string{"Hearts", "Diamond", "Clover", "Spades"}
	cardValues := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) display() {
	for i, card := range d {
		fmt.Println(i, " ", card)
	}
	fmt.Println("No.of cards ->", len(d))
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToDisk(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromDisk(fileName string) deck {
	cardsByteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	cardsStringSlice := strings.Split(string(cardsByteSlice), ",")
	return cardsStringSlice
}

func (d deck) shuffleDeck() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		randNumber := r.Intn(len(d) - 1)
		d[i], d[randNumber] = d[randNumber], d[i]
	}
	return d
}
