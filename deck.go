package main

import (
	"fmt"
	"os"
	"strings"
)

// create a new type of deck, a slice of strings
type deck []string

// No receiver here because we just need to create and not update, need to be a full deck
func newDeck() deck {
	complete := deck{}
	suites := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suite := range suites {
		for _, value := range values {
			complete = append(complete, value+" Of "+suite)
		}
	}
	return complete
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) convertToByteSlice() []byte {
	holder := strings.Join([]string(d), ",")
	byteHolder := []byte(holder)
	return byteHolder
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, d.convertToByteSlice(), 0666)
}

/*
What exactly do we plan to return here?
A deck or a part of the deck in some other form, personally, I would want to return a full
proper deck

? Read a file that does not exist --> error --> Place an error message
*/
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// 1. Log the error and send a new deck out
		// 2. Log the error and say bye
		fmt.Println("Error:", err)
		os.Exit(1) //! Code 1 for error
	}
	return deck(strings.Split(string(bs), ","))
}
