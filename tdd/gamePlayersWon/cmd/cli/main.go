package main

import (
	"bytes"
	"fmt"
	poker "goLangLearning/tdd/gamePlayersWon"
	"log"
	"os"
)

const dbFileName = "game.db.json"

var dummyStdOut = &bytes.Buffer{}

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	cli := poker.NewCLI(os.Stdin, dummyStdOut, game)

	fmt.Println("Let's play poker")
	fmt.Println("type {Name} wins to record a win")

	cli.PlayPoker()

}
