package main

import (
	"fmt"
	"goLangLearning/blackJackAI/blackjack"
)

func main() {
	fmt.Println("Blackjack AI")
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println("Number of games win:", winnings)
}
