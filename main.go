package main

// import "fmt"

func main() {
	var fullDeck deck = newDeck()
	firstHand, remainingDeck := deal(fullDeck, len(fullDeck)/2)
	firstHand.saveToFile("firstHand")
	remainingDeck.saveToFile("remainingDeck")
}

func newCard() string {
	return "Ace of Diamonds"
}
