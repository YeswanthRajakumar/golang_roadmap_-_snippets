package main

import (
// "fmt"
)

func main() {
	cards := newDeckFromFile("./cards/demo.txt")
	cards.shuffleDeck()
	cards.display()
}
