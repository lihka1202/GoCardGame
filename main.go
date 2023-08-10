package main

import "fmt"

// import "fmt"

func main() {
	var fullDeck deck = newDeck()
	fullDeck = fullDeck.shuffle()
	firstHand, remainingDeck := deal(fullDeck, len(fullDeck)/2)
	firstHand.saveToFile("firstHand")
	remainingDeck.saveToFile("remainingDeck")
	var read1, read2 deck
	read1 = newDeckFromFile("firstHand")
	read2 = newDeckFromFile("remainingDeck")
	fmt.Println(read1)
	fmt.Println(read2)
}

func newCard() string {
	return "Ace of Diamonds"
}
