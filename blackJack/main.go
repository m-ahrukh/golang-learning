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
	cards := deck.New(deck.Deck(3))
	shuffledCards := deck.Shuffle(cards)

	// fmt.Println(shuffledCards[1:6])
	// var cardsList CardsInHand = shuffledCards[1:6]
	// fmt.Println(cardsList)

	var player, dealer CardsInHand

	var card deck.Card
	for i := 0; i < 2; i++ {
		card, shuffledCards = drawCard(shuffledCards)
		player = append(player, card)
		card, shuffledCards = drawCard(shuffledCards)
		dealer = append(dealer, card)
	}

	var input string
	for input != "s" {
		fmt.Println("Player:", player)
		fmt.Println("Dealer:", dealer.DealerString())

		fmt.Println("Press h for Hit or s for Stand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, shuffledCards = drawCard(shuffledCards)
			player = append(player, card)
		}
	}

	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, shuffledCards = drawCard(shuffledCards)
		dealer = append(dealer, card)
	}

	fmt.Println("----STAND----")
	fmt.Println("Player Cards:", player, "\n---->Score is:", player.Score())
	fmt.Println("\nDealer Cards:", dealer, "\n---->Score is:", dealer.Score())
	fmt.Println()

	switch {
	case player.Score() > 21:
		fmt.Println("YOU BUSTED!")
	case dealer.Score() > 21:
		fmt.Println("DEALER BUSTED")
	case player.Score() > dealer.Score():
		fmt.Println("YOU WIN!!")
	case player.Score() < dealer.Score():
		fmt.Println("YOU LOSE!")
	case player.Score() == dealer.Score():
		fmt.Println("GAME IS DRAW")
	}
}

func drawCard(cards []deck.Card) (deck.Card, []deck.Card) {
	card := cards[0]
	remainingDeck := cards[1:]
	return card, remainingDeck
}
