package main

// import "fmt"


func main() {

	cards := newDeck() // creating a new deck of cards using the newDeck function

	// hand, remainingCards := deal(cards,5)
	// hand.print() // printing the hand of cards
	// remainingCards.print() // printing the remaining cards in the deck

	cards.saveToFile("my_cards") // saving the deck of cards to a file called "my_cards"
}
