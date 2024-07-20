package arraysandslices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	got := Sum(numbers)
	want := 21
	assert.Equal(t, want, got)
}

func TestSums(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assert.Equal(t, want, got)
	})

	t.Run("collection of 9 nummbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := Sum(numbers)
		want := 55
		assert.Equal(t, want, got)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	assert.Equal(t, want, got)
}

func TestSumAllTails(t *testing.T) {
	// got := SumAllTails([]int{}, []int{0, 5, 4})
	// want := []int{0, 9}

	// assert.Equal(t, want, got)

	t.Run("make sum of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2})
		want := []int{2}

		assert.Equal(t, want, got)
	})
}
