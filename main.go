package main

func main() {
	cards := newDeck()
	cards.print()
}

func newCard() string {
	return "Six of Spades"
}
