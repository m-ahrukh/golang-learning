//go:generate .\stringer.exe -type=Suit,Rank
package deck

import "fmt"

type Card struct {
	Suit
	Rank
}

type Rank uint8
type Suit uint8

const (
	Spade Suit = iota //iota initialized the value to 0 and then increment it in constants
	Diamond
	Club
	Heart
)

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

func (c Card) String() string {
	return fmt.Sprintf("%s of %ss\n", c.Rank.String(), c.Suit.String())
}
