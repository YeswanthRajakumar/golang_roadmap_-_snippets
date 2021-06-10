package main

// import "fmt"

func main() {
	cards := createDeck()
	cards.saveToDisk("demo.txt")
}

