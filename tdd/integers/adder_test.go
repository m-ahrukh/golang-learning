package integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	assert.Equal(t, expected, sum)
}
