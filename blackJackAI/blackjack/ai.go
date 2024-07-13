package blackjack

import (
	"fmt"
	deck "goLangLearning/deckOfCards"
)

// Reshuffling -> done
// add betting -> done
// blackjack payouts -> done
// doubing down (double the bet before dealer give card)
// splitting 7,7 (if same number of cards in one hand, split it in two bets)

type AI interface {
	Bet(shuffled bool) int
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

func (ai humanAI) Bet(shuffled bool) int {
	if shuffled {
		fmt.Println("Deck was just shuffled")
	}
	fmt.Println("What would you like to bet")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

func (ai dealerAI) Bet(shuffled bool) int {
	return 1
}

func (ai humanAI) Play(player []deck.Card, dealer deck.Card) Move {
	var input string
	for {
		fmt.Println("Player:", player)
		fmt.Println("Dealer:", dealer)

		fmt.Println("Press h for Hit, s for Stand or d for Double")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		case "d":
			return MoveDouble
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
}

func (ai dealerAI) Results(player [][]deck.Card, dealer []deck.Card) {
}
