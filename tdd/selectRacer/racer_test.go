package selectRacer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// func TestRacer(t *testing.T) {
// 	slowServer := makeDelayedServer(20 * time.Millisecond)
// 	fastServer := makeDelayedServer(0 * time.Millisecond)

// 	defer slowServer.Close()
// 	defer fastServer.Close()

// 	slowUrl := slowServer.URL
// 	fastUrl := fastServer.URL

// 	want := fastUrl
// 	got := Racer(slowUrl, fastUrl)

// 	assert.Equal(t, want, got)
// }

func TestRacers(t *testing.T) {
	t.Run("compares speed of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		assert.Equal(t, want, got)
	})

	t.Run("returns an error is a server doesn't response within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
