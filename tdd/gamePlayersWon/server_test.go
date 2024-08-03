package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {

		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}

	server := NewPlayerServer(&store)

	t.Run("it records wins on POST", func(t *testing.T) {

		player := "Pepper"

		request := newPostWinsRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q, want %q", store.winCalls[0], player)
		}

		AssertPlayerWin(t, &store, player)
	})
}

func TestLeague(t *testing.T) {

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		store := StubPlayerStore{}
		server := NewPlayerServer(&store)

		server.ServeHTTP(response, request)

		var got []Player

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response, jsonContentType)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := NewPlayerServer(&StubPlayerStore{})

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("when we get a message over a websocket it is a winner of a game", func(t *testing.T) {
		store := &StubPlayerStore{}
		winner := "Ruth"
		server := httptest.NewServer(NewPlayerServer(store))
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			t.Fatalf("could not open a ws connection on %s %v", wsURL, err)
		}
		defer ws.Close()

		if err := ws.WriteMessage(websocket.TextMessage, []byte(winner)); err != nil {
			t.Fatalf("could not send message over ws connection %v", err)
		}

		time.Sleep(10 * time.Millisecond)
		AssertPlayerWin(t, store, winner)
	})
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from srver %q into slice f Player, '%v'", body, err)
	}

	return
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostWinsRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
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

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot %v, \nwant %v", got, want)
	}
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}
