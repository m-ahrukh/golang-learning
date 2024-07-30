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
	store  PlayerStore
	router *http.ServeMux
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{
		store,
		http.NewServeMux(),
	}

	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	p.router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	return p
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/*
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

		// player := strings.TrimPrefix(r.URL.Path, "/players/")
		// switch r.Method {
		// case http.MethodPost:
		// 	p.processWin(w, player)
		// case http.MethodGet:
		// 	p.showScore(w, player)
		// }
	*/

	// router := http.NewServeMux()

	/*
		// router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 	w.WriteHeader(http.StatusOK)
		// }))

		// router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 	player := strings.TrimPrefix(r.URL.Path, "/players/")
		// 	switch r.Method {
		// 	case http.MethodPost:
		// 		p.processWin(w, player)
		// 	case http.MethodGet:
		// 		p.showScore(w, player)
		// 	}
		// }))
	*/
	// router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	// router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	// router.ServeHTTP(w, r)

	p.router.ServeHTTP(w, r)
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

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
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
