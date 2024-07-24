package propertybasedtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanNumerals(t *testing.T) {
	// got := ConvertToRoman(1)
	// want := "I"
	// assert.Equal(t, want, got)

	// t.Run("1 gets converted to I", func(t *testing.T) {
	// 	got := ConvertToRoman(1)
	// 	want := "I"

	// 	assert.Equal(t, want, got)
	// })

	// t.Run("2 gets converted to II", func(t *testing.T) {
	// 	got := ConvertToRoman(2)
	// 	want := "II"

	// 	assert.Equal(t, want, got)
	// })

	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			assert.Equal(t, test.Want, got)
		})
	}
}
