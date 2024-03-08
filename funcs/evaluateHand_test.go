package funcs

import (
	"furyoftroy/pokerbuddy/v1/objects"
	"testing"
)

func TestEvaluateHandNoCards(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("Expected error when evaluating empty hand")
		}
	}()

	EvaluateHand(make([]*objects.Card, 0))
}

func TestEvaluateHandStraightFlush(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(13, 1),
		deck.Take(12, 1),
		deck.Take(11, 1),
		deck.Take(10, 1),
		deck.Take(9, 1),
		deck.Take(8, 1))

	expectedHand := objects.NewHand(objects.STRAIGHTFLUSHRANK, cards[:5])
	checkHand(cards, expectedHand, t)

	deck.ReturnAll(cards)
}

func TestEvaluateHandFourOfAKind(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(14, 2),
		deck.Take(14, 3),
		deck.Take(14, 4),
		deck.Take(13, 1),
		deck.Take(13, 2),
		deck.Take(13, 3))

	expectedHand := objects.NewHand(objects.FOUROFAKINDRANK, cards[:5])
	checkHand(cards, expectedHand, t)

	deck.ReturnAll(cards)
}

func TestEvaluateHandFullHouse(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(14, 2),
		deck.Take(14, 3),
		deck.Take(13, 1),
		deck.Take(13, 2),
		deck.Take(12, 1),
		deck.Take(12, 2))

	expectedHand := objects.NewHand(objects.FULLHOUSERANK, cards[:5])
	checkHand(cards, expectedHand, t)

	deck.ReturnAll(cards)
}

func TestEvaluateHandFlush(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(13, 1),
		deck.Take(12, 1),
		deck.Take(11, 1),
		deck.Take(9, 1),
		deck.Take(8, 1),
		deck.Take(7, 1))

	expectedHand := objects.NewHand(objects.FLUSHRANK, cards[:5])
	checkHand(cards, expectedHand, t)

	deck.ReturnAll(cards)
}

func TestEvaluateHandStraight(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(13, 2),
		deck.Take(12, 3),
		deck.Take(11, 4),
		deck.Take(10, 1),
		deck.Take(9, 2),
		deck.Take(8, 3))

	expectedHand := objects.NewHand(objects.STRAIGHTRANK, cards[:5])
	checkHand(cards, expectedHand, t)

	deck.ReturnAll(cards)
}

func TestEvaluateHandThreeOfAKind(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(14, 2),
		deck.Take(14, 3),
		deck.Take(13, 1),
		deck.Take(12, 2),
		deck.Take(11, 3),
		deck.Take(9, 4))

	expectedHand := objects.NewHand(objects.THREEOFAKINDRANK, cards[:5])
	checkHand(cards, expectedHand, t)

	deck.ReturnAll(cards)
}

func TestEvaluateHandTwoPair(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(14, 2),
		deck.Take(13, 1),
		deck.Take(13, 2),
		deck.Take(12, 1),
		deck.Take(11, 2),
		deck.Take(9, 3))

	expectedHand := objects.NewHand(objects.TWOPAIRRANK, cards[:5])
	checkHand(cards, expectedHand, t)

	deck.ReturnAll(cards)
}

func TestEvaluateHandPair(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(14, 2),
		deck.Take(13, 1),
		deck.Take(12, 2),
		deck.Take(11, 3),
		deck.Take(9, 4),
		deck.Take(8, 1))

	expectedHand := objects.NewHand(objects.PAIRRANK, cards[:5])
	checkHand(cards, expectedHand, t)

	deck.ReturnAll(cards)
}

func TestEvaluateHandHighCard(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(13, 2),
		deck.Take(12, 3),
		deck.Take(11, 4),
		deck.Take(9, 1),
		deck.Take(8, 2),
		deck.Take(7, 3))

	expectedHand := objects.NewHand(objects.HIGHCARDRANK, cards[:5])
	checkHand(cards, expectedHand, t)

	deck.ReturnAll(cards)
}

func TestEvaluateHandStraightFlushNotStraight(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(13, 3),
		deck.Take(12, 3),
		deck.Take(11, 3),
		deck.Take(10, 3),
		deck.Take(9, 3),
		deck.Take(8, 1))

	expectedHand := objects.NewHand(objects.STRAIGHTFLUSHRANK,
		append(make([]*objects.Card, 0),
			cards[1],
			cards[2],
			cards[3],
			cards[4],
			cards[5]))
	checkHand(cards, expectedHand, t)

	deck.ReturnAll(cards)
}

func TestEvaluateHandFlushNotStraight(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(13, 2),
		deck.Take(12, 3),
		deck.Take(11, 1),
		deck.Take(10, 1),
		deck.Take(9, 1),
		deck.Take(8, 1))

	expectedHand := objects.NewHand(objects.FLUSHRANK,
		append(make([]*objects.Card, 0),
			cards[0],
			cards[3],
			cards[4],
			cards[5],
			cards[6]))
	checkHand(cards, expectedHand, t)

	deck.ReturnAll(cards)
}

func TestEvaluateHandStraightNotThreeOfAKind(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(14, 2),
		deck.Take(14, 3),
		deck.Take(13, 1),
		deck.Take(12, 2),
		deck.Take(11, 3),
		deck.Take(10, 4))

	expectedHand := objects.NewHand(objects.STRAIGHTRANK,
		append(make([]*objects.Card, 0),
			cards[0],
			cards[3],
			cards[4],
			cards[5],
			cards[6]))
	checkHand(cards, expectedHand, t)

	deck.ReturnAll(cards)
}

func TestEvaluateHandStraightRandomOrder(t *testing.T) {
	deck := objects.NewDeck()
	cards := append(make([]*objects.Card, 0),
		deck.Take(14, 1),
		deck.Take(8, 2),
		deck.Take(10, 3),
		deck.Take(7, 1),
		deck.Take(6, 2),
		deck.Take(2, 3),
		deck.Take(9, 4))

	expectedHand := objects.NewHand(objects.STRAIGHTRANK,
		append(make([]*objects.Card, 0),
			cards[2],
			cards[6],
			cards[1],
			cards[3],
			cards[4]))
	checkHand(cards, expectedHand, t)

	deck.ReturnAll(cards)
}

func checkHand(cards []*objects.Card, expectedHand *objects.Hand, t *testing.T) {
	actualHand := EvaluateHand(cards)
	if objects.CompareHands(expectedHand, actualHand) != 0 {
		t.Fatalf("Expected %s, got %s", expectedHand, actualHand)
	}
}
