package game

type Bowling struct {
	rolls       [21]int
	currentRoll int
}

func (b *Bowling) RollBall(pins int) {
	b.rolls[b.currentRoll] = pins
	b.currentRoll++
}

func (b *Bowling) Score() int {
	score := 0
	indexInFrame := 0
	for frameIndex := 0; frameIndex < 10; frameIndex++ {
		if b.isStrike(indexInFrame) {
			score += 10 + b.strikeBonus(indexInFrame)
			indexInFrame++
		} else if b.isSpare(indexInFrame) {
			score += 10 + b.spareBonus(indexInFrame)
			indexInFrame += 2
		} else {
			score += b.rolls[indexInFrame] + b.rolls[indexInFrame+1]
			indexInFrame += 2
		}
	}
	return score
}

func (b *Bowling) isSpare(indexInFrame int) bool {
	return b.rolls[indexInFrame]+b.rolls[indexInFrame+1] == 10
}

func (b *Bowling) isStrike(indexInFrame int) bool {
	return b.rolls[indexInFrame] == 10
}

func (b *Bowling) strikeBonus(indexInFrame int) int {
	return b.rolls[indexInFrame+1] + b.rolls[indexInFrame+2]
}

func (b *Bowling) spareBonus(indexInFrame int) int {
	return b.rolls[indexInFrame+2]
}
