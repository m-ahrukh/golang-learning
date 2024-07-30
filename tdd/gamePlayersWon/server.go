package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayersScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// if r.Method == http.MethodPost {
	// 	w.WriteHeader(http.StatusAccepted)
	// 	return
	// }

	// player := strings.TrimPrefix(r.URL.Path, "/players/")
	// score := p.store.GetPlayersScore(player)
	// if score == 0 {
	// 	w.WriteHeader(http.StatusNotFound)
	// }
	// fmt.Fprint(w, score)

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

	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score := p.store.GetPlayersScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	// p.store.RecordWin("Bob")
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
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
