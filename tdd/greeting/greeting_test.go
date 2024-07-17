package greeting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		// given
		name := "Aamir"
		// when
		got := SayHello(name)
		// then
		assert.Equal(t, "Hello, Aamir!", got)
	})

	t.Run("say hello to the world", func(t *testing.T) {
		// when
		got := SayHello("")
		// then
		assert.Equal(t, "Hello, World!", got)
	})
}

/*
func TestHelloWithName(t *testing.T) {
	// given
	name := "Aamir"
	// when
	got := SayHello(name)
	// then
	assert.Equal(t, "Hello, Aamir!", got)
}

func TestHelloWithoutName(t *testing.T) {
	// when
	got := SayHello("")
	// then
	assert.Equal(t, "Hello, World!", got)
}
*/
