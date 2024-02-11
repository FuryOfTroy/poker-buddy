package funcs

import (
	"fmt"
	"furyoftroy/pokerfriend/v1/objects"
	"strings"
)

func MergeHandMaps(m1 map[int][]*objects.PossibleHand, m2 map[int][]*objects.PossibleHand) {
	for name, possibleHands := range m2 {
		m1[name] = append(m1[name], possibleHands...)
	}
}

func CalculateHandOdds(cards []*objects.Card, deck *objects.Deck) map[int][]*objects.PossibleHand {
	results := make(map[int][]*objects.PossibleHand)
	resultChan := make(chan map[int][]*objects.PossibleHand)

	resultCount := 0
	for i := 0; i < 52; i++ {
		if EnumerateRemainingCardsInGoroutine(i, cards, deck, resultChan) {
			resultCount++
		}
	}

	totalHands := 1
	for i := len(cards); i < 7; i++ {
		totalHands = totalHands * (52 - i)
	}
	fmt.Printf("\nCalculating about %d hands...\n\n", totalHands)

	for i := 0; i < resultCount; i++ {
		chanResults := <-resultChan
		MergeHandMaps(results, chanResults)
	}

	return results
}

func EnumerateRemainingCardsInGoroutine(i int, _cards []*objects.Card, _deck *objects.Deck, resultChan chan map[int][]*objects.PossibleHand) bool {
	deck := _deck.Clone()
	card := deck.TryTakeIndex(i)
	if card != nil {
		cards := make([]*objects.Card, len(_cards))
		copy(cards, _cards)
		go InternalEnumerateRemainingCardsInGoroutine(append(cards, card), deck, 1, resultChan)
		return true
	}
	return false
}

func InternalEnumerateRemainingCardsInGoroutine(cards []*objects.Card, deck *objects.Deck, outCount int, resultChan chan map[int][]*objects.PossibleHand) {
	handsByRank := make(map[int][]*objects.PossibleHand)
	defer func() {
		resultChan <- handsByRank
	}()
	EnumerateRemainingCards(cards, deck, outCount, handsByRank)
}

func EnumerateRemainingCards(cards []*objects.Card, deck *objects.Deck, outCount int, handsByRank map[int][]*objects.PossibleHand) {
	if len(cards) == 7 {
		hand := EvaluateHand(cards)
		if hand == nil {
			ReportNoHandCalculated(cards)
		} else {
			outs := make([]*objects.Card, 0)
			for i := len(cards) - outCount; i < len(cards); i++ {
				outs = append(outs, cards[i])
			}
			handsByRank[hand.GetRank()] = append(handsByRank[hand.GetRank()], objects.NewPossibleHand(hand, outs))
			return
		}
		panic(fmt.Errorf("Uh oh, should have returned or panicked"))
	}

	for i := 0; i < 52; i++ {
		card := deck.TryTakeIndex(i)
		if card == nil {
			continue
		}
		func() {
			defer deck.Return(card)
			EnumerateRemainingCards(append(cards, card), deck, outCount+1, handsByRank)
		}()
	}
}

func ReportNoHandCalculated(cards []*objects.Card) {
	var b strings.Builder
	b.WriteString("No hand produced! Cards: ")
	for _, card := range cards {
		b.WriteString("|")
		b.WriteString(card.Print())
		b.WriteString("|")
	}
	panic(fmt.Errorf(b.String()))
}
