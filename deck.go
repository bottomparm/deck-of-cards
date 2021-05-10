package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cardSuits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardVals := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	var newDeck deck
	for _, suit := range cardSuits {
		for _, val := range cardVals {
			card := val + " of " + suit
			newDeck = append(newDeck, card)
		}
	}
	return newDeck
}

// any variable of type deck now has access to the print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString () string {
	stringSlice := []string(d)
	return strings.Join(stringSlice, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1: Log the error and return a call to newDeck().
		// Option 2: Log the error and exit the program completely.
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	size := len(d) - 1
	for i := range d {
		randCardIdx := r.Intn(size)
		d[i], d[randCardIdx] = d[randCardIdx], d[i]
	}
}