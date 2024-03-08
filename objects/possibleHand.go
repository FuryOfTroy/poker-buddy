package objects

import (
	"fmt"
	"strings"
)

type PossibleHand struct {
	hand *Hand
	outs []*Card
}

func NewPossibleHand(hand *Hand, outs []*Card) *PossibleHand {
	return &PossibleHand{hand: hand, outs: outs}
}

func (ph *PossibleHand) GetHand() *Hand {
	return ph.hand
}

func (ph *PossibleHand) GetOuts() []*Card {
	return ph.outs
}

func (ph *PossibleHand) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Hand: %s -- ", ph.hand))
	b.WriteString("[")
	for _, card := range ph.outs {
		b.WriteString(card.String())
	}
	b.WriteString("]")

	return b.String()
}
