package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.dealToString()), 0666)
}

func newDeckFromFile(filename string, exitOnError bool) deck {
	bs, err := os.ReadFile(filename)
	
	if err != nil {
		fmt.Println("Error: ", err)
		if exitOnError {
			os.Exit(1)
		}

		return newDeck()
	}
	return deck(strings.Split(string(bs), ", "))
}

func (d deck) shuffleDeck() {
	// Implementation for shuffling the deck
	
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for i := range d {
		newPosition := rng.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
		// fmt.Println(i, len(d), newPosition)

		println("")
	}
}