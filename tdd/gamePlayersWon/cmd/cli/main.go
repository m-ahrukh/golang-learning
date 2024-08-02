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

	fmt.Println("Let's play poker")
	fmt.Println("type {Name} wins to record a win")
	poker.NewCLI(store, os.Stdin, dummyStdOut, poker.BlindAlerterFunc(poker.StdOutAlerter)).PlayPoker()

}
