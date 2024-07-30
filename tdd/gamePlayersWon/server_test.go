package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	wincalls []string
}

func (s *StubPlayerStore) GetPlayersScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.wincalls = append(s.wincalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
	}
	// server := &PlayerServer{&store}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		// request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		// response := httptest.NewRecorder()

		// PlayerServer(response, request)

		// got := response.Body.String()
		// want := "20"

		// assert.Equal(t, want, got)

		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		// PlayerServer(response, request)
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		// request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		// response := httptest.NewRecorder()

		// PlayerServer(response, request)

		// got := response.Body.String()
		// want := "10"

		// assert.Equal(t, want, got)

		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		// PlayerServer(response, request)
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}

	// server := &PlayerServer{&store}
	server := NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("it records wins on POST", func(t *testing.T) {

		player := "Pepper"

		request := newPostWinsRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.wincalls) != 2 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.wincalls), 2)
		}

		if store.wincalls[0] != player {
			t.Errorf("did not store correct winner got %q, want %q", store.wincalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	//server := &PlayerServer{&store}
	// server := NewPlayerServer(&store)
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func newPostWinsRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
