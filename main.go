package main

import (
	"fmt"
	"furyoftroy/pokerfriend/v1/funcs"
	"furyoftroy/pokerfriend/v1/objects"
)

func printAllHandOdds(handsByRank map[int][]*objects.PossibleHand) {
	allHandsCount := 0.0
	for _, hands := range handsByRank {
		allHandsCount += float64(len(hands))
	}

	fmt.Println("---Hand Odds---")
	printHandOdds(objects.HIGHCARDRANK, handsByRank, allHandsCount)
	printHandOdds(objects.PAIRRANK, handsByRank, allHandsCount)
	printHandOdds(objects.TWOPAIRRANK, handsByRank, allHandsCount)
	printHandOdds(objects.THREEOFAKINDRANK, handsByRank, allHandsCount)
	printHandOdds(objects.STRAIGHTRANK, handsByRank, allHandsCount)
	printHandOdds(objects.FLUSHRANK, handsByRank, allHandsCount)
	printHandOdds(objects.FULLHOUSERANK, handsByRank, allHandsCount)
	printHandOdds(objects.FOUROFAKINDRANK, handsByRank, allHandsCount)
	printHandOdds(objects.STRAIGHTFLUSHRANK, handsByRank, allHandsCount)
}

func printHandOdds(rank int, handsByRank map[int][]*objects.PossibleHand, allHandsCount float64) {
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

	printAllHandOdds(possibleHandsByRank)

	deck.ReturnAll(cards)
}
