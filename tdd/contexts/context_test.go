package contexts

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SpyStore struct {
	response string
	// cancelled bool
	t *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

// func (s *SpyStore) Fetch() string {
// 	time.Sleep(100 * time.Millisecond)
// 	return s.response
// }

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Panicln("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// func (s *SpyStore) Cancel() {
// 	s.cancelled = true
// }

// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Error("store was not told to cancel")
// 	}
// }

// func (s *SpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Error("store was told to cancel")
// 	}
// }

// func TestServer(t *testing.T) {
// 	data := "hello, world"
// 	svr := Server(&SpyStore{data})

// 	request := httptest.NewRequest(http.MethodGet, "/", nil)
// 	response := httptest.NewRecorder()

// 	svr.ServeHTTP(response, request)

// 	got := response.Body.String()

// 	assert.Equal(t, data, got)
// }

func TestServer(t *testing.T) {
	// data := "hello, world"
	// t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
	// 	// data := "hello, world"
	// 	// store := &SpyStore{response: data}
	// 	store := &SpyStore{response: data, t: t}
	// 	svr := Server(store)

	// 	request := httptest.NewRequest(http.MethodGet, "/", nil)

	// 	cancellingCtx, cancel := context.WithCancel(request.Context())
	// 	time.AfterFunc(5*time.Millisecond, cancel)
	// 	request = request.WithContext(cancellingCtx)

	// 	response := httptest.NewRecorder()
	// 	svr.ServeHTTP(response, request)

	// 	if !store.cancelled {
	// 		t.Error("store was not told to cancel")
	// 	}
	// })

	// t.Run("returns data from store", func(t *testing.T) {
	// 	// data := "hello, world"
	// 	// store := &SpyStore{response: data}
	// 	store := &SpyStore{response: data, t: t}
	// 	svr := Server(store)

	// 	request := httptest.NewRequest(http.MethodGet, "/", nil)
	// 	response := httptest.NewRecorder()

	// 	svr.ServeHTTP(response, request)
	// 	got := response.Body.String()

	// 	assert.Equal(t, got, data)

	// 	if store.cancelled {
	// 		t.Error("it should not hve cancelled the store")
	// 	}
	// })

	t.Run("resturns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		got := response.Body.String()
		assert.Equal(t, data, got)
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := &SpyResponseWriter{}
		svr.ServeHTTP(response, request)
		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
