package funcs

import (
	"furyoftroy/pokerbuddy/v1/objects"
	"testing"
)

func TestCalculateHandOddsStraightFlush(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(13, 1),
		deck.Take(12, 1),
		deck.Take(11, 1),
		deck.Take(10, 1))

	hand := EvaluateHand(cards)
	handsByRank := CalculateHandOdds(hand, cards, deck)
	checkOdds(handsByRank, handsByRank, t)

	deck.ReturnAll(cards)
}

func TestCalculateHandOddsFourOfAKind(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(14, 2),
		deck.Take(14, 3),
		deck.Take(14, 4),
		deck.Take(13, 1))

	hand := EvaluateHand(cards)
	handsByRank := CalculateHandOdds(hand, cards, deck)
	checkOdds(handsByRank, handsByRank, t)

	deck.ReturnAll(cards)
}

func TestCalculateHandOddsFullHouse(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(14, 2),
		deck.Take(14, 3),
		deck.Take(13, 1),
		deck.Take(13, 2))

	hand := EvaluateHand(cards)
	handsByRank := CalculateHandOdds(hand, cards, deck)
	checkOdds(handsByRank, handsByRank, t)

	deck.ReturnAll(cards)
}

func TestCalculateHandOddsFlush(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(13, 1),
		deck.Take(12, 1),
		deck.Take(11, 1),
		deck.Take(9, 1))

	hand := EvaluateHand(cards)
	handsByRank := CalculateHandOdds(hand, cards, deck)
	checkOdds(handsByRank, handsByRank, t)

	deck.ReturnAll(cards)
}

func TestCalculateHandOddsStraight(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(13, 2),
		deck.Take(12, 3),
		deck.Take(11, 4),
		deck.Take(10, 1))

	hand := EvaluateHand(cards)
	handsByRank := CalculateHandOdds(hand, cards, deck)
	checkOdds(handsByRank, handsByRank, t)

	deck.ReturnAll(cards)
}

func TestCalculateHandOddsThreeOfAKind(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(14, 2),
		deck.Take(14, 3),
		deck.Take(13, 1),
		deck.Take(12, 2))

	hand := EvaluateHand(cards)
	handsByRank := CalculateHandOdds(hand, cards, deck)
	checkOdds(handsByRank, handsByRank, t)

	deck.ReturnAll(cards)
}

func TestCalculateHandOddsTwoPair(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(14, 2),
		deck.Take(13, 1),
		deck.Take(13, 2),
		deck.Take(12, 1))

	hand := EvaluateHand(cards)
	handsByRank := CalculateHandOdds(hand, cards, deck)
	checkOdds(handsByRank, handsByRank, t)

	deck.ReturnAll(cards)
}

func TestCalculateHandOddsPair(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(14, 2),
		deck.Take(13, 1),
		deck.Take(12, 2),
		deck.Take(11, 3))

	hand := EvaluateHand(cards)
	handsByRank := CalculateHandOdds(hand, cards, deck)
	checkOdds(handsByRank, handsByRank, t)

	deck.ReturnAll(cards)
}

func TestCalculateHandOddsHighCard(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(13, 2),
		deck.Take(12, 3),
		deck.Take(11, 4),
		deck.Take(9, 1))

	hand := EvaluateHand(cards)
	handsByRank := CalculateHandOdds(hand, cards, deck)
	checkOdds(handsByRank, handsByRank, t)

	deck.ReturnAll(cards)
}

func checkOdds(actualPossibleHandsByRank map[int][]*objects.PossibleHand, expectedPossibleHandsByRank map[int][]*objects.PossibleHand, t *testing.T) {
	// TODO: Actually check odds
	// t.Fatalf("Expected %s, got %s", expectedHand.Print(), actualHand.Print())
}
