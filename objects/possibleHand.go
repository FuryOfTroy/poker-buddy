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

func (ph *PossibleHand) Print() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Hand:%s -- ", ph.hand.Print()))
	b.WriteString("[")
	for _, card := range ph.outs {
		b.WriteString(fmt.Sprintf("%5s", card.Print()))
	}
	b.WriteString(" ]")

	return b.String()
}
