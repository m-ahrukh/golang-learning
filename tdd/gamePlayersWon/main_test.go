package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrieveingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	// server := PlayerServer{store}
	// server := NewPlayerServer(store)
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinsRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinsRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinsRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}
