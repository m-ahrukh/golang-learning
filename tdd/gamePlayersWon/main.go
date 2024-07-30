package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayersScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func main() {
	// handler := http.HandlerFunc(PlayerServer)

	// server := &PlayerServer{}
	server := &PlayerServer{NewInMemoryPlayerStore()}
	// log.Fatal(http.ListenAndServe(":3000", handler))
	log.Fatal(http.ListenAndServe(":3000", server))
}
