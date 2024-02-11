package main

import (
	"fmt"
	"furyoftroy/pokerfriend/v1/funcs"
	"furyoftroy/pokerfriend/v1/objects"
)

func PrintAllHandOdds(handsByRank map[int][]*objects.PossibleHand) {
	allHandsCount := 0.0
	for _, hands := range handsByRank {
		allHandsCount += float64(len(hands))
	}

	fmt.Println("---Hand Odds---")
	PrintHandOdds(objects.HIGHCARDRANK, handsByRank, allHandsCount)
	PrintHandOdds(objects.PAIRRANK, handsByRank, allHandsCount)
	PrintHandOdds(objects.TWOPAIRRANK, handsByRank, allHandsCount)
	PrintHandOdds(objects.THREEOFAKINDRANK, handsByRank, allHandsCount)
	PrintHandOdds(objects.STRAIGHTRANK, handsByRank, allHandsCount)
	PrintHandOdds(objects.FLUSHRANK, handsByRank, allHandsCount)
	PrintHandOdds(objects.FULLHOUSERANK, handsByRank, allHandsCount)
	PrintHandOdds(objects.FOUROFAKINDRANK, handsByRank, allHandsCount)
	PrintHandOdds(objects.STRAIGHTFLUSHRANK, handsByRank, allHandsCount)
}

func PrintHandOdds(rank int, handsByRank map[int][]*objects.PossibleHand, allHandsCount float64) {
	fmt.Printf("%-15s: %%%5.2f (%d)\n", objects.HANDNAMESBYRANK[rank], (float64(len(handsByRank[rank]))/allHandsCount)*100, len(handsByRank[rank]))
}

func main() {
	deck := objects.NewDeck()
	cards := make([]*objects.Card, 0)
	cards = append(cards,
		deck.Take(13, 1),
		deck.Take(10, 1),
		deck.Take(9, 2),
		deck.Take(4, 2),
		deck.Take(13, 2))

	if len(cards) >= 5 {
		currentHand := funcs.EvaluateHand(cards)
		fmt.Printf("Current Hand: %s\n", currentHand.Print())
	} else {
		fmt.Print("Not enough cards for current hand calculation\n")
	}
	possibleHandsByRank := funcs.CalculateHandOdds(cards, deck)

	PrintAllHandOdds(possibleHandsByRank)

	deck.ReturnAll(cards)
}
