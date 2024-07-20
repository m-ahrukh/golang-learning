package iteration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	assert.Equal(t, expected, repeated)
}

func TestRepeatNTimes(t *testing.T) {

	t.Run("repeat character 3 times", func(t *testing.T) {
		repeated := RepeatChar("a", 3)
		want := "aaa"
		assert.Equal(t, repeated, want)
	})

	t.Run("give negative number", func(t *testing.T) {
		repeated := RepeatChar("a", -1)
		want := "Number of times must be positive"
		assert.Equal(t, repeated, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
