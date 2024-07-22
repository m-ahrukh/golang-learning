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
		}
		// assert.Equal(t, err.Error(), want)
		assert.Error(t, err)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"
	// dictionary.Add("test", "this is just a test")
	dictionary.Add(word, definition)

	want := "this is just a test"
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assert.Equal(t, want, got)
}

func TestAddDefinitions(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "Test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		if err != nil {
			assert.Error(t, err)
		}

		got, err := dictionary.Search(word)
		want := "this is just a test"
		if err != nil {
			t.Fatal("should find added word", err)
		}
		assert.Equal(t, want, got)
	})

	t.Run("existing word", func(t *testing.T) {

		word := "Test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new definition")

		if err != nil {
			assert.Error(t, err)
		}

		got, err := dictionary.Search(word)
		want := "this is just a test"
		if err != nil {
			t.Fatal("should find added word", err)
		}
		assert.Equal(t, want, got)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	newDefinition := "New definition"

	dictionary.Update(word, newDefinition)

	got, err := dictionary.Search(word)
	want := newDefinition
	if err != nil {
		t.Fatal("should find added word", err)
	}
	assert.Equal(t, want, got)
}

func TestUpdates(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {

		word := "Test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := " new definition"
		err := dictionary.Update(word, newDefinition)

		if err != nil {
			assert.Error(t, err)
		}

		got, err := dictionary.Search(word)
		want := newDefinition
		if err != nil {
			t.Fatal("should find added word", err)
		}
		assert.Equal(t, want, got)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Update(word, definition)
		assert.Equal(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	// if err != nil {
	// 	assert.Error(t, err)
	// }

	assert.Equal(t, err, ErrNotFound)
}
