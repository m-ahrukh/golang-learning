package clockface

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(tm)

	assert.Equal(t, want, got)
}

// func TestSecondHandAt30Seconds(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

// 	want := Point{X: 150, Y: 150 + 90}
// 	got := SecondHand(tm)

// 	assert.Equal(t, want, got)
// }

func TestSecondsInRadians(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := secondsInRadians(thirtySeconds)

	assert.Equal(t, want, got)
}
