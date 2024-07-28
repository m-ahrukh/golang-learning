package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayersScore(name string) int {
	return 123
}

func main() {
	// handler := http.HandlerFunc(PlayerServer)

	// server := &PlayerServer{}
	server := &PlayerServer{&InMemoryPlayerStore{}}
	// log.Fatal(http.ListenAndServe(":3000", handler))
	log.Fatal(http.ListenAndServe(":3000", server))
}
