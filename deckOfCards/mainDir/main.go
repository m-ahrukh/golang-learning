package main

import (
	"fmt"
	deck "goLangLearning/deckOfCards"
)

func main() {
	cards := deck.New(deck.DefaultSort)
	fmt.Println("Total Cards: ", len(cards))
	// for _, card := range cards {
	// 	fmt.Println(card)
	// }

	shuffeledCards := deck.Shuffle(cards)
	for _, card := range shuffeledCards {
		fmt.Println(card)
	}

}
