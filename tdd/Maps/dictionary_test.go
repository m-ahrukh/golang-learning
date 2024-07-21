package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestSearch(t *testing.T) {
// 	dictionary := Dictionary{"test": "This is just a test"}

// 	got := dictionary.Search("test")
// 	want := "This is just a test"

// 	assert.Equal(t, want, got)
// }

func TestSearches(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assert.Equal(t, want, got)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		// want := "could not fint the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
			// assert.Error(t, err)
		}
		// assert.Equal(t, err.Error(), want)
		assert.Error(t, err)
	})
}
