package main

import (
	"fmt"
	deck "goLangLearning/deckOfCards"
	"strings"
)

type CardsInHand []deck.Card

func (cards CardsInHand) String() string {
	listOfCards := make([]string, len(cards))

	for index, _ := range cards {
		listOfCards[index] = cards[index].String()
	}

	return strings.Join(listOfCards, ", ")
}

func (cards CardsInHand) DealerString() string {
	return cards[0].String() + ", **HIDDEN**"
}

func (cards CardsInHand) Score() int {
	minScore := cards.MinScore()

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

func (cards CardsInHand) MinScore() int {
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

func main() {
	// cards := deck.New(deck.Deck(3))
	// shuffledCards := deck.Shuffle(cards)

	var gs GameState
	gs = Shuffle(gs)
	fmt.Println(len(gs.Deck))

	// var player, dealer CardsInHand
	// var card deck.Card
	// for i := 0; i < 2; i++ {
	// 	card, shuffledCards = drawCard(shuffledCards)
	// 	player = append(player, card)
	// 	card, shuffledCards = drawCard(shuffledCards)
	// 	dealer = append(dealer, card)
	// }

	gs = Deal(gs)

	// var input string
	// for input != "s" {
	// 	fmt.Println("Player:", player)
	// 	fmt.Println("Dealer:", dealer.DealerString())

	// 	fmt.Println("Press h for Hit or s for Stand")
	// 	fmt.Scanf("%s\n", &input)
	// 	switch input {
	// 	case "h":
	// 		card, shuffledCards = drawCard(shuffledCards)
	// 		player = append(player, card)
	// 	}
	// }

	var input string
	for gs.State == StatePlayerTurn {
		fmt.Println("Player:", gs.Player)
		fmt.Println("Dealer:", gs.Dealer.DealerString())

		fmt.Println("Press h for Hit or s for Stand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			// card, shuffledCards = drawCard(shuffledCards)
			// player = append(player, card)
			gs = Hit(gs)
		case "s":
			gs = Stand(gs)
		default:
			fmt.Println("Invalid Option:", input)
		}
	}

	// for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
	// 	card, shuffledCards = drawCard(shuffledCards)
	// 	dealer = append(dealer, card)
	// }

	for gs.State == StateDealerTurn {
		if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
			gs = Hit(gs)
		} else {
			gs = Stand(gs)
		}
	}

	// fmt.Println("----STAND----")
	// fmt.Println("Player Cards:", player, "\n---->Score is:", player.Score())
	// fmt.Println("\nDealer Cards:", dealer, "\n---->Score is:", dealer.Score())
	// fmt.Println()

	// switch {
	// case player.Score() > 21:
	// 	fmt.Println("YOU BUSTED!")
	// case dealer.Score() > 21:
	// 	fmt.Println("DEALER BUSTED")
	// case player.Score() > dealer.Score():
	// 	fmt.Println("YOU WIN!!")
	// case player.Score() < dealer.Score():
	// 	fmt.Println("YOU LOSE!")
	// case player.Score() == dealer.Score():
	// 	fmt.Println("GAME IS DRAW")
	// }

	gs = EndGame(gs)
}

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Deck(3))
	ret.Deck = deck.Shuffle(ret.Deck)
	return ret
}

func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(CardsInHand, 0, 10)
	ret.Dealer = make(CardsInHand, 0, 10)

	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = drawCard(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = drawCard(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}
	ret.State = StatePlayerTurn
	return ret
}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	currentPlayer := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = drawCard(ret.Deck)
	*currentPlayer = append(*currentPlayer, card)

	if currentPlayer.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

func EndGame(gs GameState) GameState {
	ret := clone(gs)
	fmt.Println("----STAND----")
	fmt.Println("Player Cards:", ret.Player, "\n---->Score is:", ret.Player.Score())
	fmt.Println("\nDealer Cards:", ret.Dealer, "\n---->Score is:", ret.Dealer.Score())
	fmt.Println()

	switch {
	case ret.Player.Score() > 21:
		fmt.Println("YOU BUSTED!")
	case ret.Dealer.Score() > 21:
		fmt.Println("DEALER BUSTED")
	case ret.Player.Score() > ret.Dealer.Score():
		fmt.Println("YOU WIN!!")
	case ret.Player.Score() < ret.Dealer.Score():
		fmt.Println("YOU LOSE!")
	case ret.Player.Score() == ret.Dealer.Score():
		fmt.Println("GAME IS DRAW")
	}
	ret.Player = nil
	ret.Dealer = nil
	return ret
}

func drawCard(cards []deck.Card) (deck.Card, []deck.Card) {
	card := cards[0]
	remainingDeck := cards[1:]
	return card, remainingDeck
}

// To simplify game for AI
type GameState struct {
	Deck   []deck.Card
	State  State
	Player CardsInHand
	Dealer CardsInHand
}

type State uint8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

func (gs *GameState) CurrentPlayer() *CardsInHand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("It is not currently any player's turn")
	}
}

func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(CardsInHand, len(gs.Player)),
		Dealer: make(CardsInHand, len(gs.Dealer)),
	}

	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}
