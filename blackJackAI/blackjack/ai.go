package blackjack

import (
	"fmt"
	deck "goLangLearning/deckOfCards"
)

type AI interface {
	Bet() int
	Play(player []deck.Card, dealer deck.Card) Move
	Results(player [][]deck.Card, dealer []deck.Card)
}

type humanAI struct {
}

type dealerAI struct {
}

func HumanAI() AI {
	return humanAI{}
}

func (ai humanAI) Bet() int {
	return 1
}

func (ai dealerAI) Bet() int {
	return 1
}

func (ai humanAI) Play(player []deck.Card, dealer deck.Card) Move {
	var input string
	for {
		fmt.Println("Player:", player)
		fmt.Println("Dealer:", dealer)

		fmt.Println("Press h for Hit or s for Stand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		default:
			fmt.Println("Invalid Option:", input)
		}
	}
}

func (ai dealerAI) Play(player []deck.Card, dealer deck.Card) Move {
	if Score(player...) <= 16 || (Score(player...) == 17 && Soft(player...)) {
		return MoveHit
	} else {
		return MoveStand
	}
}

func (ai humanAI) Results(player [][]deck.Card, dealer []deck.Card) {
	fmt.Println("----FINAL HAND----")
	fmt.Println("Player Cards:", player)
	fmt.Println("Dealer Cards:", dealer)
	fmt.Println("----------------------------------------")
}

func (ai dealerAI) Results(player [][]deck.Card, dealer []deck.Card) {
}
