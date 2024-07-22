package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Mahrukh")

	got := buffer.String()
	want := "Hello, Mahrukh"

	assert.Equal(t, want, got)
}
