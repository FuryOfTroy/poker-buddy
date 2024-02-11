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

func (o *PossibleHand) Print() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Hand: %s -- ", o.hand.Print()))
	b.WriteString("[")
	for _, card := range o.outs {
		b.WriteString(fmt.Sprintf("|%s|", card.Print()))
	}
	b.WriteString("]")

	return b.String()
}
