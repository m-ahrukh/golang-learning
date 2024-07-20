package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	assert.Equal(t, want, got)
}
