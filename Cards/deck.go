package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}                                                // creating an empty deck of cards
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} // defining the suits of the cards
	cardValues := []string{"Ace", "Two", "Three", "Four"}          // defining the values of the cards

	for _, suit := range cardSuits { // iterating over the suits of the cards
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit) // appending a new card to the deck by combining the value and suit of the card
		}
	}
	return cards
}

func (d deck) print() { // defining a method called print on the deck type which takes a deck as a receiver and prints the values in the deck
	for i, card := range d {
		fmt.Println(i, card) // printing the index and value of each card in the deck
	}
}

func deal(d deck, handSize int) (deck, deck) { // defining a function called deal which takes a deck and a hand size as arguments and returns two decks - one for the hand and one for the remaining cards
	return d[:handSize], d[handSize:] // returning the hand and the remaining cards by slicing the original deck
}

func (d deck) toString() string { // defining a method called toString on the deck type which takes a deck as a receiver and returns a string representation of the deck
	return strings.Join([]string(d), ",") // converting the deck to a slice of strings and joining them with a comma to create a single string representation of the deck
}

func (d deck) saveToFile(filename string) error { // defining a method called saveToFile on the deck type which takes a deck as a receiver and a filename as an argument and saves the deck to a file
	return os.WriteFile(filename, []byte(d.toString()), 0666) // writing the string representation of the deck to a file with the specified filename and permissions
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) // reading the contents of the file with the specified filename
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ",")) // converting the file contents to a slice of strings and creating a new deck
}

func (d deck) shuffle() { // defining a method called shuffle on the deck type which takes a deck as a receiver and shuffles the cards in the deck
	source := rand.NewSource(time.Now().UnixNano()) // creating a new random source using the process ID to ensure different shuffling each time
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)           // generating a random index for the new position of the card
		d[i], d[newPosition] = d[newPosition], d[i] // swapping the current card with the card at the new position
	}
}
