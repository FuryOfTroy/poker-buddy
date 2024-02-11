package funcs

import (
	"fmt"
	"furyoftroy/pokerfriend/v1/objects"
	"sort"
)

func GetCardsByValue(cards []*objects.Card) map[int][]*objects.Card {
	cardsByValue := make(map[int][]*objects.Card)
	for _, card := range cards {
		cardsByValue[card.GetValue()] = append(cardsByValue[card.GetValue()], card)
	}
	return cardsByValue
}

func GetCardGroupsByCount(cardsByValue map[int][]*objects.Card) map[int][][]*objects.Card {
	cardGroupsByCount := make(map[int][][]*objects.Card)

	sortedValues := GetValuesDesc(cardsByValue)
	for _, value := range sortedValues {
		cardGroup := cardsByValue[value]
		cardGroupsByCount[len(cardGroup)] = append(cardGroupsByCount[len(cardGroup)], cardGroup)
	}
	return cardGroupsByCount
}

func GetValuesDesc(_map map[int][]*objects.Card) []int {
	keys := make([]int, 0)
	for key := range _map {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	return keys
}

func EvaluateHand(_cards []*objects.Card) *objects.Hand {
	if len(_cards) < 5 {
		panic(fmt.Errorf("Cannot Eval a hand with less than 5 cards"))
	}
	cards := make([]*objects.Card, len(_cards))
	copy(cards, _cards)

	sort.Slice(cards, func(i int, j int) bool {
		return cards[i].GetValue() > cards[j].GetValue()
	})

	flushOrStraightFlush := getFlushOrStraightFlush(cards)
	if flushOrStraightFlush != nil {
		return flushOrStraightFlush
	}

	cardsByValue := GetCardsByValue(cards)
	cardGroupsByCount := GetCardGroupsByCount(cardsByValue)

	//Check 4-of-a-kind
	fourOfAKinds := cardGroupsByCount[4]
	if fourOfAKinds != nil {
		fourOfAKind := fourOfAKinds[0]
		sortedValues := GetValuesDesc(cardsByValue)
		if sortedValues[0] != fourOfAKind[0].GetValue() {
			return objects.NewHand(objects.FOUROFAKINDRANK, append(fourOfAKind, cardsByValue[sortedValues[0]]...))
		}
		return objects.NewHand(objects.FOUROFAKINDRANK, append(fourOfAKind, cardsByValue[sortedValues[1]]...))
	}

	//Check 3-of-a-kind and Full House
	if cardGroupsByCount[3] != nil {
		threeOfAKind := cardGroupsByCount[3][0]
		if cardGroupsByCount[2] != nil {
			return objects.NewHand(objects.FULLHOUSERANK, append(threeOfAKind, cardGroupsByCount[2][0]...))
		}

		straight := getStraight(cardsByValue)
		if straight != nil {
			return objects.NewHand(objects.STRAIGHTRANK, straight)
		}

		kickers := make([]*objects.Card, 0)
		for _, card := range cards {
			if card.GetValue() != threeOfAKind[0].GetValue() {
				kickers = append(kickers, card)
				if len(kickers) == 2 {
					return objects.NewHand(objects.THREEOFAKINDRANK, append(threeOfAKind, kickers...))
				}
			}
		}

		panic(fmt.Errorf("Couldn't get kickers for 3-of-a-kind"))
	}

	straight := getStraight(cardsByValue)
	if straight != nil {
		return objects.NewHand(objects.STRAIGHTRANK, straight)
	}

	// Check for pairs
	twoOfAKinds := cardGroupsByCount[2]
	if twoOfAKinds != nil {
		if len(twoOfAKinds) >= 2 {
			for _, card := range cards {
				if card.GetValue() != twoOfAKinds[0][0].GetValue() && card.GetValue() != twoOfAKinds[1][0].GetValue() {
					twoPairs := append(twoOfAKinds[0], twoOfAKinds[1]...)
					return objects.NewHand(objects.TWOPAIRRANK, append(twoPairs, card))
				}
			}

			panic(fmt.Errorf("Couldn't get kickers for 2-pair"))
		}

		pair := twoOfAKinds[0]
		kickers := make([]*objects.Card, 0)
		for _, card := range cards {
			if card.GetValue() != pair[0].GetValue() {
				kickers = append(kickers, card)
				if len(kickers) == 3 {
					return objects.NewHand(objects.PAIRRANK, append(pair, kickers...))
				}
			}
		}
		panic(fmt.Errorf("Couldn't get kickers for pair"))
	}

	return objects.NewHand(objects.HIGHCARDRANK, cards[:5])
}

// isStraightFlush checks if the given slice of Cards forms a straight flush.
func getFlushOrStraightFlush(cards []*objects.Card) *objects.Hand {
	if cards == nil {
		return nil
	}

	flushCards := getFlush(cards)
	if flushCards != nil {
		straightFlushCards := getStraight(GetCardsByValue(flushCards))
		if straightFlushCards != nil {
			return objects.NewHand(objects.STRAIGHTFLUSHRANK, straightFlushCards)
		}
		return objects.NewHand(objects.FLUSHRANK, flushCards[:5])
	}

	return nil
}

// isFlush checks if the given slice of Cards forms a flush.
func getFlush(cards []*objects.Card) []*objects.Card {
	if cards == nil {
		return nil
	}

	cardsBySuit := make(map[int][]*objects.Card)
	for _, card := range cards {
		cardsBySuit[card.GetSuit()] = append(cardsBySuit[card.GetSuit()], card)
	}

	for _, cardsGroup := range cardsBySuit {
		if len(cardsGroup) >= 5 {
			return cardsGroup
		}
	}

	return nil
}

// isFlush checks if the given slice of Cards forms a flush.
func getStraight(cardsByValue map[int][]*objects.Card) []*objects.Card {
	if cardsByValue == nil {
		return nil
	}

	lowCard := 99
	highCard := 99

	sortedValues := GetValuesDesc(cardsByValue)
	if sortedValues[0] == 14 {
		sortedValues = append(sortedValues, 1)
	}
	for _, value := range sortedValues {
		// TODO: Treat a 14 (Ace) as 1, so maybe lowCard-value == -12
		if lowCard-value != 1 {
			highCard = value
		} else if highCard-value >= 4 {
			if value == 1 {
				return append(make([]*objects.Card, 0), cardsByValue[5][0],
					cardsByValue[4][0],
					cardsByValue[3][0],
					cardsByValue[2][0],
					cardsByValue[14][0])
			} else {
				return append(make([]*objects.Card, 0), cardsByValue[highCard][0],
					cardsByValue[highCard-1][0],
					cardsByValue[highCard-2][0],
					cardsByValue[highCard-3][0],
					cardsByValue[highCard-4][0])
			}
		}
		lowCard = value
	}

	return nil
}
