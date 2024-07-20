package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
//function without any parameter
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	assert.Equal(t, want, got)
}
*/

/*
// Test with parameter
func TestHello(t *testing.T) {
	got := Hello("Mahrukh")
	want := "Hello, Mahrukh"

	assert.Equal(t, want, got)
}
*/

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Mahrukh", "")
		want := "Hello, Mahrukh"

		assertCorrectMessage(t, want, got)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, want, got)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Mahrukh", "Spanish")
		want := "Hola, Mahrukh"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Mahrukh", "French")
		want := "Bonjour, Mahrukh"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	assert.Equal(t, got, want)
}
