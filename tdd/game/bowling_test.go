package game_test

import (
	"goLangLearning/tdd/game"
	"testing"

	"github.com/stretchr/testify/assert"
)

func rollMany(game *game.Bowling, rolls, pins int) {
	for i := 0; i < rolls; i++ {
		game.RollBall(pins)
	}
}

func rollSpare(game *game.Bowling) {
	game.RollBall(5)
	game.RollBall(5)
}

func rollStrike(game *game.Bowling) {
	game.RollBall(10)
}

func TestRollBall(t *testing.T) {
	game := &game.Bowling{}

	game.RollBall(1)

	assert.Equal(t, 1, game.Score())
}

func TestScore(t *testing.T) {
	game := &game.Bowling{}

	game.RollBall(5)

	assert.Equal(t, 5, game.Score())
}

func TestGutterGame(t *testing.T) {
	game := &game.Bowling{}

	rollMany(game, 20, 0)

	assert.Equal(t, 0, game.Score())
}

func TestAllOnes(t *testing.T) {
	game := &game.Bowling{}

	rollMany(game, 20, 1)

	assert.Equal(t, 20, game.Score())
}

func TestOneSpare(t *testing.T) {
	game := &game.Bowling{}

	rollSpare(game)
	game.RollBall(3)
	rollMany(game, 17, 0)

	assert.Equal(t, 16, game.Score())
}

func TestOneStrike(t *testing.T) {
	game := &game.Bowling{}

	rollStrike(game)

	game.RollBall(3)
	game.RollBall(4)
	rollMany(game, 16, 0)

	assert.Equal(t, 24, game.Score())
}

func TestPerfectGame(t *testing.T) {
	game := &game.Bowling{}

	rollMany(game, 12, 10)

	assert.Equal(t, 300, game.Score())
}
