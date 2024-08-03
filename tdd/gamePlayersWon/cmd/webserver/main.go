package main

import (
	poker "goLangLearning/tdd/gamePlayersWon"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := mustMakePlayerServer(store)

	log.Fatal(http.ListenAndServe(":3000", server))
}

func mustMakePlayerServer(store poker.PlayerStore) *poker.PlayerServer {
	server, err := poker.NewPlayerServer(store)
	if err != nil {
		log.Fatal("problem creating player server", err)
	}
	return server
}
