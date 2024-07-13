package main

import (
	"fmt"
	"goLangLearning/blackJackAI/blackjack"
	deck "goLangLearning/deckOfCards"
)

type basicAI struct {
	score int
	seen  int
	decks int
}

func (ai *basicAI) Bet(shuffled bool) int {
	if shuffled {
		ai.score = 0
		ai.seen = 0
	}
	// fmt.Println("true score:", (ai.decks*52-ai.seen)/52)
	trueScore := ai.score / ((ai.decks * 52) / 52)
	switch {
	case trueScore >= 14:
		return 100000
	case trueScore >= 8:
		return 5000
	default:
		return 100
	}
}

func (ai *basicAI) Play(hand []deck.Card, dealer deck.Card) blackjack.Move {
	score := blackjack.Score(hand...)
	if len(hand) == 2 {
		if hand[0] == hand[1] {
			cardScore := blackjack.Score(hand[0])
			if cardScore >= 8 && cardScore != 10 {
				return blackjack.MoveSplit
			}
		}
		if (score == 10 || score == 11) && !blackjack.Soft(hand...) {
			return blackjack.MoveDouble
		}
	}
	dealerScore := blackjack.Score(dealer)
	if dealerScore >= 5 && dealerScore <= 6 {
		return blackjack.MoveStand
	}
	if score < 13 {
		return blackjack.MoveHit
	}
	return blackjack.MoveStand
}

func (ai *basicAI) Results(player [][]deck.Card, dealer []deck.Card) {
	for _, card := range dealer {
		ai.count(card)
	}
	for _, hand := range player {
		for _, card := range hand {
			ai.count(card)
		}
	}
}

func (ai *basicAI) count(card deck.Card) {
	score := blackjack.Score(card)
	switch {
	case score >= 10:
		ai.score--
	case score <= 6:
		ai.score++
	}
	ai.seen++
}

func main() {
	fmt.Println("Blackjack AI")
	option := blackjack.Options{
		Decks:           4,
		Hands:           20,
		BlackjackPayout: 1.4,
	}
	game := blackjack.New(option)
	winnings := game.Play(&basicAI{
		decks: 4,
	})
	fmt.Println("Your Score is:", winnings)
}
