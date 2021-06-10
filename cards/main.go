package main

func main() {
	cards := createDeck()
	// cards.display()
	cardsInHand , remainingDeck := deal(cards,5)
	cardsInHand.display()
	remainingDeck.display()

}

