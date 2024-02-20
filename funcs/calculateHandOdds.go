package funcs

import (
	"fmt"
	"furyoftroy/pokerbuddy/v1/objects"
	"strings"
)

func CalculateHandOdds(originalHand *objects.Hand, cards []*objects.Card, deck *objects.Deck) map[int][]*objects.PossibleHand {
	results := make(map[int][]*objects.PossibleHand)
	if len(cards) == 7 {
		return results
	}

	resultChan := make(chan map[int][]*objects.PossibleHand)

	resultCount := 0
	for i := 51; i >= 0; i-- {
		if enumerateRemainingCardsInGoroutine(i, originalHand, cards, deck, resultChan) {
			resultCount++
		}
	}

	totalHands := 1
	for i := len(cards); i < 7; i++ {
		totalHands = totalHands * (52 - i)
	}
	fmt.Printf("\nAnalyzing about %d possible hands...\n\n", totalHands)

	for i := 0; i < resultCount; i++ {
		chanResults := <-resultChan
		mergeHandMaps(results, chanResults)
	}

	return results
}

func mergeHandMaps(m1 map[int][]*objects.PossibleHand, m2 map[int][]*objects.PossibleHand) {
	for name, possibleHands := range m2 {
		m1[name] = append(m1[name], possibleHands...)
	}
}

func enumerateRemainingCardsInGoroutine(i int, originalHand *objects.Hand, _cards []*objects.Card, _deck *objects.Deck, resultChan chan map[int][]*objects.PossibleHand) bool {
	deck := _deck.Clone()
	card := deck.TryTakeIndex(i)
	if card != nil {
		cards := make([]*objects.Card, len(_cards))
		copy(cards, _cards)
		go internalEnumerateRemainingCardsInGoroutine(originalHand, append(cards, card), deck, 1, resultChan)
		return true
	}
	return false
}

func internalEnumerateRemainingCardsInGoroutine(originalHand *objects.Hand, cards []*objects.Card, deck *objects.Deck, outCount int, resultChan chan map[int][]*objects.PossibleHand) {
	handsByRank := make(map[int][]*objects.PossibleHand)
	defer func() {
		resultChan <- handsByRank
	}()
	enumerateRemainingCards(originalHand, cards, deck, outCount, handsByRank)
}

func enumerateRemainingCards(originalHand *objects.Hand, cards []*objects.Card, deck *objects.Deck, outCount int, handsByRank map[int][]*objects.PossibleHand) {
	handToBeat := originalHand
	if len(cards) >= 5 {
		hand := EvaluateHand(cards)
		if hand == nil {
			reportNoHandCalculated(cards)
		} else {
			if objects.CompareHands(hand, originalHand) > 0 {
				handToBeat = hand
				outs := make([]*objects.Card, 0)
				for i := len(cards) - outCount; i < len(cards); i++ {
					for _, cardInHand := range hand.GetCards() {
						if cards[i] == cardInHand {
							outs = append(outs, cards[i])
						}
					}
				}
				handsByRank[hand.GetRank()] = append(handsByRank[hand.GetRank()], objects.NewPossibleHand(hand, outs))
			}
		}
	}

	if len(cards) < 7 {
		lastCard := cards[len(cards)-1]
		for i := objects.GetIndex(lastCard); i >= 0; i-- {
			card := deck.TryTakeIndex(i)
			if card == nil {
				continue
			}
			func() {
				defer deck.Return(card)
				enumerateRemainingCards(handToBeat, append(cards, card), deck, outCount+1, handsByRank)
			}()
		}
	}
}

func reportNoHandCalculated(cards []*objects.Card) {
	var b strings.Builder
	b.WriteString("No hand produced! Cards: ")
	for _, card := range cards {
		b.WriteString(fmt.Sprintf("|%-4s|", card.Print()))
	}
	panic(fmt.Errorf(b.String()))
}
