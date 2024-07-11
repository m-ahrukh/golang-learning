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

	fmt.Println("Stand")
	fmt.Println("Player Cards:", player)
	fmt.Println("Dealer Cards:", dealer)
}

func drawCard(cards []deck.Card) (deck.Card, []deck.Card) {
	card := cards[0]
	remainingDeck := cards[1:]
	return card, remainingDeck
}
