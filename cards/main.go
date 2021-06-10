package main

import (
	"fmt"
)

func main() {
	// cardsString,err := newDeckFromFile("demo.txt")
	cards := newDeckFromFile("./cards/demo.txt")
	fmt.Println(cards)
}