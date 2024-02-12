package main

import (
	"fmt"
	"furyoftroy/pokerfriend/v1/funcs"
	"furyoftroy/pokerfriend/v1/objects"
)

func printAllHandStats(handsByRank map[int][]*objects.PossibleHand) {
	allHandsCount := 0.0
	for _, hands := range handsByRank {
		allHandsCount += float64(len(hands))
	}

	fmt.Println("---Hand Odds---")
	printHandStats(objects.HIGHCARDRANK, handsByRank, allHandsCount)
	printHandStats(objects.PAIRRANK, handsByRank, allHandsCount)
	printHandStats(objects.TWOPAIRRANK, handsByRank, allHandsCount)
	printHandStats(objects.THREEOFAKINDRANK, handsByRank, allHandsCount)
	printHandStats(objects.STRAIGHTRANK, handsByRank, allHandsCount)
	printHandStats(objects.FLUSHRANK, handsByRank, allHandsCount)
	printHandStats(objects.FULLHOUSERANK, handsByRank, allHandsCount)
	printHandStats(objects.FOUROFAKINDRANK, handsByRank, allHandsCount)
	printHandStats(objects.STRAIGHTFLUSHRANK, handsByRank, allHandsCount)
}

func printHandStats(rank int, handsByRank map[int][]*objects.PossibleHand, allHandsCount float64) {
	printHandOdds(rank, handsByRank, allHandsCount)
	printHandsAndOuts(handsByRank[rank])
}

func printHandsAndOuts(possibleHands []*objects.PossibleHand) {
	fmt.Print("\tPossible hands:\n")
	if len(possibleHands) < 10 {
		for _, ph := range possibleHands {
			fmt.Printf("\t%s\n", ph.Print())
		}
	} else {
		for i := 0; i < 5; i++ {
			fmt.Printf("\t%s\n", possibleHands[i].Print())
		}
		fmt.Println("...")
		for i := len(possibleHands) - 6; i < len(possibleHands)-1; i++ {
			fmt.Printf("\t%s\n", possibleHands[i].Print())
		}
	}
}

func printHandOdds(rank int, handsByRank map[int][]*objects.PossibleHand, allHandsCount float64) {
	fmt.Printf("%-15s: %5.2f%% (%d)\n", objects.HANDNAMESBYRANK[rank], (float64(len(handsByRank[rank]))/allHandsCount)*100, len(handsByRank[rank]))
}

func main() {
	deck := objects.NewDeck()
	cards := make([]*objects.Card, 0)
	cards = append(cards,
		deck.Take(13, 1),
		deck.Take(10, 1),
		deck.Take(7, 2),
		deck.Take(9, 2),
		deck.Take(5, 2))

	if len(cards) >= 5 {
		currentHand := funcs.EvaluateHand(cards)
		fmt.Printf("Current Hand: %s\n", currentHand.Print())
	} else {
		fmt.Print("Not enough cards for current hand calculation\n")
	}
	possibleHandsByRank := funcs.CalculateHandOdds(cards, deck)

	printAllHandStats(possibleHandsByRank)

	deck.ReturnAll(cards)
}
