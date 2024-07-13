package main

import (
	"fmt"
	"goLangLearning/blackJackAI/blackjack"
)

func main() {
	fmt.Println("Blackjack AI")
	option := blackjack.Options{
		Decks:           2,
		Hands:           2,
		BlackjackPayout: 1.4,
	}
	game := blackjack.New(option)
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
