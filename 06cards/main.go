package main

import "fmt"

func main() {
	cards := deck{newCard(), "New card added"}
	fmt.Println("Main function", cards)

	cards = append(cards, "One more card")

	for i, card := range cards {
		fmt.Println("Card is ", i, card)
	}
}

func newCard() string {

	return "Five of Diamonds"
}
