package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "")
	defer cleanDatabase()
	store, err := NewFileSystemPlayerStore(database)
	AssertNoError(t, err)
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), NewPostWinsRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinsRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinsRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewGetScoreRequest(player))
		AssertStatus(t, response.Code, http.StatusOK)

		AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewLeagueRequest())
		AssertStatus(t, response.Code, http.StatusOK)

		got := GetLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		AssertLeague(t, got, want)
	})
}
