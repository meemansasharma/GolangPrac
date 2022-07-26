package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
	cards.shuffleDeck()
	fmt.Println("Shuffled cards")
	cards.print()

}
