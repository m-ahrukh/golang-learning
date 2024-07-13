package blackjack

import (
	"errors"
	"fmt"
	deck "goLangLearning/deckOfCards"
)

type state int8

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

type Options struct {
	Decks           int
	Hands           int
	BlackjackPayout float64
}

type Game struct {
	// unexported fields
	deck            []deck.Card
	state           state
	player          []hand
	playerBet       int
	handIdx         int
	dealer          []deck.Card
	dealerAI        AI
	balance         int
	nDecks          int
	nHands          int
	blackjackPayout float64
}

type hand struct {
	cards []deck.Card
	bet   int
}

func New(option Options) Game {
	g := Game{
		state:    statePlayerTurn,
		dealerAI: &dealerAI{},
		balance:  0,
	}

	if option.Decks == 0 {
		option.Decks = 3
	}
	if option.Hands == 0 {
		option.Hands = 100
	}
	if option.BlackjackPayout == 0.0 {
		option.BlackjackPayout = 1.5
	}
	g.nDecks = option.Decks
	g.nHands = option.Hands
	g.blackjackPayout = option.BlackjackPayout

	return g
}

func bet(g *Game, ai AI, shuffled bool) {
	bet := ai.Bet(shuffled)
	if bet < 100 {
		panic("bet must be atleast 100")
	}
	g.playerBet = bet
}

func deal(g *Game) {
	playerHand := make([]deck.Card, 0, 10)
	g.handIdx = 0
	g.dealer = make([]deck.Card, 0, 10)

	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = drawCard(g.deck)
		playerHand = append(playerHand, card)
		card, g.deck = drawCard(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.player = []hand{
		{
			cards: playerHand,
			bet:   g.playerBet,
		},
	}
	g.state = statePlayerTurn
}

func Blackjack(hand ...deck.Card) bool {
	return len(hand) == 2 && Score(hand...) == 21
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

	dealerScore := Score(g.dealer...)
	dealerBlackjack := Blackjack(g.dealer...)
	allHands := make([][]deck.Card, len(g.player))

	for hIndex, hand := range g.player {
		cards := hand.cards
		allHands[hIndex] = cards
		playerScore := Score(cards...)
		playerBlackjack := Blackjack(cards...)
		winnings := hand.bet

		switch {
		case playerBlackjack && dealerBlackjack:
			fmt.Println("GAME IS DRAW")
			winnings = 0
		case dealerBlackjack:
			fmt.Println("DEALER'S BLACKJACK")
			winnings *= -1
		case playerBlackjack:
			fmt.Println("PLAYER'S BLACKJACK")
			winnings = int(g.blackjackPayout * float64(winnings))
		case playerScore > 21:
			fmt.Println("YOU BUSTED!")
			// g.balance--
			winnings *= -1
		case dealerScore > 21:
			fmt.Println("DEALER BUSTED")

			// g.balance++
		case playerScore > dealerScore:
			fmt.Println("YOU WIN!!")
			// g.balance++
		case playerScore < dealerScore:
			fmt.Println("YOU LOSE!")
			// g.balance--
			winnings *= -1
		case playerScore == dealerScore:
			fmt.Println("GAME IS DRAW")
			winnings = 0
		}
		fmt.Println()
		g.balance += winnings
	}

	ai.Results(allHands, g.dealer)
	// fmt.Println("Score is:", Score(g.player...))
	// fmt.Println("Score is:", Score(g.dealer...))
	// fmt.Println("----------------------------------------")

	// playerBlackjack, dealerBlackjack := Blackjack(g.player...), Blackjack(g.dealer...)

	// winnings := g.playerBet
	// switch {
	// case playerBlackjack && dealerBlackjack:
	// 	fmt.Println("GAME IS DRAW")
	// 	winnings = 0
	// case dealerBlackjack:
	// 	fmt.Println("DEALER'S BLACKJACK")
	// 	winnings *= -1
	// case playerBlackjack:
	// 	fmt.Println("PLAYER'S BLACKJACK")
	// 	winnings = int(g.blackjackPayout * float64(winnings))
	// case Score(g.player...) > 21:
	// 	fmt.Println("YOU BUSTED!")
	// 	// g.balance--
	// 	winnings *= -1
	// case Score(g.dealer...) > 21:
	// 	fmt.Println("DEALER BUSTED")

	// 	// g.balance++
	// case Score(g.player...) > Score(g.dealer...):
	// 	fmt.Println("YOU WIN!!")
	// 	// g.balance++
	// case Score(g.player...) < Score(g.dealer...):
	// 	fmt.Println("YOU LOSE!")
	// 	// g.balance--
	// 	winnings *= -1
	// case Score(g.player...) == Score(g.dealer...):
	// 	fmt.Println("GAME IS DRAW")
	// 	winnings = 0
	// }
	// fmt.Println()
	// g.balance += winnings

	g.player = nil
	g.dealer = nil
}

func (g *Game) Play(ai AI) int {
	// g.deck = deck.New(deck.Deck(g.nDecks))
	// g.deck = deck.Shuffle(g.deck)

	g.deck = nil

	//for shuffling -> shufflilng condition
	min := 52 * g.nDecks / 5

	for i := 0; i < g.nHands; i++ {

		shuffled := false
		if len(g.deck) < min {
			g.deck = deck.New(deck.Deck(g.nDecks))
			g.deck = deck.Shuffle(g.deck)
			shuffled = true
		}

		//betting stage
		bet(g, ai, shuffled)

		deal(g)

		if Blackjack(g.dealer...) {
			endGame(g, ai)
			continue
		}

		for g.state == statePlayerTurn {
			// player := make([]deck.Card, len(g.player))
			// copy(player, g.player)
			hand := make([]deck.Card, len(*g.currentPlayer()))
			copy(hand, *g.currentPlayer())
			move := ai.Play(hand, g.dealer[0])
			err := move(g)
			switch err {
			case errBust:
				MoveStand(g)
			case nil:
			default:
				panic(err)
			}
		}

		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}

		endGame(g, ai)
	}
	return g.balance
}

type Move func(*Game) error

var (
	errBust = errors.New("hand score exceeded 21")
)

func (g *Game) currentPlayer() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player[g.handIdx].cards
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("It is not currently any player's turn")
	}
}

func MoveHit(g *Game) error {

	currentPlayer := g.currentPlayer()
	var card deck.Card
	card, g.deck = drawCard(g.deck)
	*currentPlayer = append(*currentPlayer, card)

	if Score(*currentPlayer...) > 21 {
		// MoveStand(g)
		return errBust
	}
	return nil
}

func MoveDouble(g *Game) error {
	if len(*g.currentPlayer()) != 2 {
		return errors.New("can only double on a hand of two cards")

	}
	g.playerBet *= 2
	MoveHit(g)
	return MoveStand(g)
}

func MoveStand(g *Game) error {
	if g.state == stateDealerTurn {
		g.state++
		return nil
	}
	if g.state == statePlayerTurn {
		g.handIdx++
		if g.handIdx >= len(g.player) {
			g.state++
		}
		return nil

	}
	return errors.New("invalid case")
}

func MoveSplit(g *Game) error {
	cards := g.currentPlayer()
	if len(*cards) != 2 {
		return errors.New("you can only split with two of the same cards")
	}
	if (*cards)[0].Rank != (*cards)[1].Rank {
		return errors.New("both cards must be of the same rank")
	}
	g.player = append(g.player, hand{
		cards: []deck.Card{(*cards)[1]},
		bet:   g.player[g.handIdx].bet,
	})
	g.player[g.handIdx].cards = (*cards)[:1]
	return nil
}

func drawCard(cards []deck.Card) (deck.Card, []deck.Card) {
	card := cards[0]
	remainingDeck := cards[1:]
	return card, remainingDeck
}
