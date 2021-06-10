package main

import "fmt"

// Something like class deck extends string slice
type deck []string

func createDeck() deck{
	cards := deck{}
	cardSuites := [4]string{"Hearts","Diamond","Clover","Spade"}
	cardValues := [13]string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jack","Queen","King"}
	
	for _,suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

func (d deck) display()  {
	for i, card := range d {
		fmt.Println(i," ",card)
	}
	fmt.Println("No.of cards ->",len(d))
}

func deal(d deck, handSize int) (deck,deck) {
	return d[:handSize],d[handSize:]
	
}