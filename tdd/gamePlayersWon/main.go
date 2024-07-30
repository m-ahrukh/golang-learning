package main

import (
	"log"
	"net/http"
)

func main() {
	// handler := http.HandlerFunc(PlayerServer)

	// server := &PlayerServer{}
	server := &PlayerServer{NewInMemoryPlayerStore()}
	// log.Fatal(http.ListenAndServe(":3000", handler))
	log.Fatal(http.ListenAndServe(":3000", server))
}
