package main

// decsize := 10 -- assignment outside of the function - not allowed
func main() {
	// fmt.Println(decsize) - shows erroer
	// card := newcard()
	// fmt.Println(card)

	// cards := []string{"Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades", newcard()}
	// // fmt.Println(cards) // works as expected return the elements in the slice with a space separator

	// cards = append(cards, "Six of Spades") // appending a new element to the slice
	// for i, card := range cards {
	// 	fmt.Println(i, card) // prints the index and the card
	// }
	// cards := deck{"Ace of Spades", newcard()}
	// cards = append(cards, "Two of Spades", "Three of Spades", "Four of Spades")
	//cards := newdeck() // creating a new deck of cards
	// cards.print()
	// hand, remainingDeck := deal(cards, 5)
	// fmt.Println("The hand")
	// hand.print()
	// fmt.Println("The remaining deck")
	// remainingDeck.print() // dealing 5 cards from the deck

	//byte scilicing
	// greeting := "Hello, World!"
	// fmt.Println([]byte(greeting)) // prints the greeting message
	//fmt.Println(cards.toString())
	// cards.saveToFile("mycards")

	newCardsFromFiles := newdeckFromFile("mycards")
	newCardsFromFiles.print()   // reading the cards from the file and printing them
	newCardsFromFiles.shuffle() // shuffling the cards
	newCardsFromFiles.print()   // printing the shuffled cards
}

// func newcard() string {
// 	return "Five of Diamonds"
// }
