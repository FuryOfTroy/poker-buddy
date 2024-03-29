package objects

import (
	"fmt"
	"strings"
)

const HIGHCARDNAME = "High Card"
const PAIRNAME = "Pair"
const TWOPAIRNAME = "2-pair"
const THREEOFAKINDNAME = "3-of-a-kind"
const STRAIGHTNAME = "Straight"
const FLUSHNAME = "Flush"
const FULLHOUSENAME = "Full House"
const FOUROFAKINDNAME = "4-of-a-kind"
const STRAIGHTFLUSHNAME = "StraightFlush"

const HIGHCARDRANK = 0
const PAIRRANK = 1
const TWOPAIRRANK = 2
const THREEOFAKINDRANK = 3
const STRAIGHTRANK = 4
const FLUSHRANK = 5
const FULLHOUSERANK = 6
const FOUROFAKINDRANK = 7
const STRAIGHTFLUSHRANK = 8

var HANDNAMESBYRANK = map[int]string{
	HIGHCARDRANK:      HIGHCARDNAME,
	PAIRRANK:          PAIRNAME,
	TWOPAIRRANK:       TWOPAIRNAME,
	THREEOFAKINDRANK:  THREEOFAKINDNAME,
	STRAIGHTRANK:      STRAIGHTNAME,
	FLUSHRANK:         FLUSHNAME,
	FULLHOUSERANK:     FULLHOUSENAME,
	FOUROFAKINDRANK:   FOUROFAKINDNAME,
	STRAIGHTFLUSHRANK: STRAIGHTFLUSHNAME}

type Hand struct {
	rank  int
	cards []*Card
}

func NewHand(rank int, cards []*Card) *Hand {
	if rank < 0 || rank > 8 {
		panic(fmt.Errorf("Hand cannot have rank of %d", rank))
	}
	return &Hand{rank: rank, cards: cards}
}

func CompareHands(h1 *Hand, h2 *Hand) int {
	if h1 == h2 {
		return 0
	}

	if h1 == nil {
		return -1
	}

	if h2 == nil {
		return 1
	}

	if h1.rank > h2.rank {
		return 1
	}

	if h1.rank < h2.rank {
		return -1
	}

	for i := 0; i < 5; i++ {
		if h1.cards[i].value > h2.cards[i].value {
			return 1
		}
		if h1.cards[i].value < h2.cards[i].value {
			return -1
		}
	}

	return 0
}

func (h *Hand) GetRank() int {
	return h.rank
}

func (h *Hand) GetName() string {
	return HANDNAMESBYRANK[h.rank]
}

func (h *Hand) GetFullName() string {
	valueDesc := "N/A"
	switch h.rank {
	case 0:
		valueDesc = h.cards[0].GetValueString()
	case 1:
		valueDesc = h.cards[0].GetValueString()
	case 2:
		valueDesc = fmt.Sprintf("%s/%s", h.cards[0].GetValueString(), h.cards[2].GetValueString())
	case 3:
		valueDesc = h.cards[0].GetValueString()
	case 4:
		valueDesc = h.cards[0].GetValueString()
	case 5:
		valueDesc = h.cards[0].String()
	case 6:
		valueDesc = fmt.Sprintf("%s/%s", h.cards[0].GetValueString(), h.cards[3].GetValueString())
	case 7:
		valueDesc = h.cards[0].GetValueString()
	case 8:
		valueDesc = h.cards[0].String()
	}

	return fmt.Sprintf("%s{%s}", h.GetName(), valueDesc)
}

func (h *Hand) GetCards() []*Card {
	return h.cards
}

func (h *Hand) String() string {
	var b strings.Builder
	for _, card := range h.cards {
		b.WriteString(card.String())
	}
	b.WriteString(fmt.Sprintf("(%s)", h.GetFullName()))

	return b.String()
}
