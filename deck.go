package main

import (
	"fmt"
	"strings"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}

	fmt.Println("Length of deck: ", len(d))

}

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	
	return cards;
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// func dealToString(d deck, handSize int) string {
// 	return strings.Join(d[:handSize], ", ")
// }

func (d deck) dealToString() string {
	return strings.Join(d, ", ")
}