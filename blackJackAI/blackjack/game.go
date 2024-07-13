package blackjack

import (
	"fmt"
	deck "goLangLearning/deckOfCards"
)

type state int8

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

type Game struct {
	// unexported fields
	deck     []deck.Card
	state    state
	player   []deck.Card
	dealer   []deck.Card
	dealerAI AI
	balance  int
}

func New() Game {
	return Game{
		state:    statePlayerTurn,
		dealerAI: &dealerAI{},
		balance:  0,
	}
}

func Deal(g *Game) {
	g.player = make([]deck.Card, 0, 10)
	g.dealer = make([]deck.Card, 0, 10)

	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = drawCard(g.deck)
		g.player = append(g.player, card)
		card, g.deck = drawCard(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.state = statePlayerTurn
}

func Score(cards ...deck.Card) int {
	minScore := minScore(cards...)

	if minScore > 11 {
		return minScore
	}
	for _, card := range cards {
		//currently ace value is 1. if minScore is less than 11, change the value of Ace from 1 to 11
		if card.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore

}

func Soft(cards ...deck.Card) bool {
	minScore := minScore(cards...)
	score := Score(cards...)
	// if minScore != score {
	// 	return true
	// }
	// return false

	return minScore != score
}

func minScore(cards ...deck.Card) int {
	score := 0
	for _, card := range cards {
		score += min(int(card.Rank), 10)
	}
	return score
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func endGame(g *Game, ai AI) {
	fmt.Println("Score is:", Score(g.player...))
	fmt.Println("Score is:", Score(g.dealer...))
	switch {
	case Score(g.player...) > 21:
		fmt.Println("YOU BUSTED!")
		g.balance--
	case Score(g.dealer...) > 21:
		fmt.Println("DEALER BUSTED")
		g.balance++
	case Score(g.player...) > Score(g.dealer...):
		fmt.Println("YOU WIN!!")
		g.balance++
	case Score(g.player...) < Score(g.dealer...):
		fmt.Println("YOU LOSE!")
		g.balance--
	case Score(g.player...) == Score(g.dealer...):
		fmt.Println("GAME IS DRAW")
	}
	fmt.Println()
	ai.Results([][]deck.Card{g.player}, g.dealer)
	fmt.Println()
	g.player = nil
	g.dealer = nil
}

func (g *Game) Play(ai AI) int {
	g.deck = deck.New(deck.Deck(3))
	g.deck = deck.Shuffle(g.deck)

	for i := 0; i < 2; i++ {

		Deal(g)

		for g.state == statePlayerTurn {
			player := make([]deck.Card, len(g.player))
			copy(player, g.player)
			move := ai.Play(player, g.dealer[0])
			move(g)
		}

		for g.state == stateDealerTurn {
			dealer := make([]deck.Card, len(g.dealer))
			copy(dealer, g.dealer)
			move := g.dealerAI.Play(dealer, g.dealer[0])
			move(g)
		}

		endGame(g, ai)
	}
	return g.balance
}

type Move func(*Game)

func (g *Game) currentPlayer() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("It is not currently any player's turn")
	}
}

func MoveHit(g *Game) {

	currentPlayer := g.currentPlayer()
	var card deck.Card
	card, g.deck = drawCard(g.deck)
	*currentPlayer = append(*currentPlayer, card)

	if Score(*currentPlayer...) > 21 {
		MoveStand(g)
	}

	// return
}

func MoveStand(g *Game) {
	g.state++
}

func drawCard(cards []deck.Card) (deck.Card, []deck.Card) {
	card := cards[0]
	remainingDeck := cards[1:]
	return card, remainingDeck
}
