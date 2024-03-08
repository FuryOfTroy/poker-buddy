package main

import (
	"fmt"
	"furyoftroy/pokerbuddy/v1/app"
	"furyoftroy/pokerbuddy/v1/funcs"
	"furyoftroy/pokerbuddy/v1/objects"
	"log"
	"os"

	"github.com/spf13/cobra"
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
	if len(possibleHands) > 0 {
		fmt.Printf("\t%s\n", possibleHands[0])
	}
}

func printHandOdds(rank int, handsByRank map[int][]*objects.PossibleHand, allHandsCount float64) {
	fmt.Printf("%-15s: %5.2f%% (%d)\n", objects.HANDNAMESBYRANK[rank], (float64(len(handsByRank[rank]))/allHandsCount)*100, len(handsByRank[rank]))
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "pokerbuddy",
		Short: "pokerbuddy - a CLI that helps you play poker",
		Long: `pokerbuddy is a CLI that takes information about the state of your poker game, and provides useful statistics that can guide your next decision
	   
	You might not be allowed to use this in online play, so use with caution!`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	var webCmd = &cobra.Command{
		Use:   "web",
		Short: "Launch the pokerbuddy web app",
		Long:  "Launch the pokerbuddy web app so you can use a GUI instead of the CLI",
		Run: func(cmd *cobra.Command, args []string) {
			host, found := os.LookupEnv("PB_HOST")
			if !found {
				host = "localhost"
			}
			port, found := os.LookupEnv("PB_PORT")
			if !found {
				port = "41100"
			}

			app := app.NewApplication()
			log.Fatal(app.Listen(fmt.Sprintf("%s:%s", host, port)))
		},
	}

	var evalCmd = &cobra.Command{
		Use:   "eval",
		Short: "Evaluate your current hand",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			eval(args[0])
		},
	}

	rootCmd.AddCommand(webCmd)
	rootCmd.AddCommand(evalCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func eval(cardsStr string) {
	deck := objects.NewDeck()
	cards := make([]*objects.Card, 0)
	for i := 0; i < len(cardsStr); i += 2 {
		cards = append(cards, deck.TakeName(cardsStr[i:i+2]))
	}

	var currentHand *objects.Hand
	if len(cards) >= 5 {
		currentHand = funcs.EvaluateHand(cards)
		fmt.Printf("Current Hand: %s\n", currentHand)
	} else {
		fmt.Print("Not enough cards for current hand calculation\n")
	}
	possibleHandsByRank := funcs.CalculateHandOdds(currentHand, cards, deck)

	printAllHandStats(possibleHandsByRank)

	deck.ReturnAll(cards)
}
