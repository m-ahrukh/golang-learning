package main

import (
	"fmt"
	poker "goLangLearning/tdd/gamePlayersWon"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewTexesHoldem(poker.BlindAlerterFunc(poker.Alerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)

	fmt.Println("Let's play poker")
	fmt.Println("type {Name} wins to record a win")

	cli.PlayPoker()

}
