package main

func main() {
	cards := deck{"Ace","Diamonds","clover","Hearts"}
	cards = append(cards, newCard())
	cards.display()

}

func newCard() string{
	return "5 of Spades"
}