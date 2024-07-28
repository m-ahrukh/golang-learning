package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayersScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.store.GetPlayersScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	// w.WriteHeader(http.StatusNotFound)

	// if player == "Pepper" {
	// 	fmt.Fprint(w, "20")
	// 	return
	// }

	// if player == "Floyd" {
	// 	fmt.Fprint(w, "10")
	// 	return
	// }

	// fmt.Fprint(w, p.store.GetPlayersScore(player))
	fmt.Fprint(w, score)

}

// func GetPlayersScore(name string) string {
// 	if name == "Pepper" {
// 		return "20"
// 	}
// 	if name == "Floyd" {
// 		return "10"
// 	}
// 	return ""
// }
