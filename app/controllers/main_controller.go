package controllers

import (
	"encoding/json"
	"fmt"
	"furyoftroy/pokerbuddy/v1/app/serialization"
	"furyoftroy/pokerbuddy/v1/funcs"
	"furyoftroy/pokerbuddy/v1/objects"

	"github.com/gofiber/fiber/v2"
)

func RenderIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func EvaluateCards(c *fiber.Ctx) error {
	deck := objects.NewDeck()
	requestBody := c.Context().Request.Body()
	var cardNames []string
	err := json.Unmarshal(requestBody, &cardNames)
	if err != nil {
		panic(err)
	}
	cards := make([]*objects.Card, 0)
	for _, cardName := range cardNames {
		cards = append(cards, deck.TakeName(cardName))
	}
	hand := funcs.EvaluateHand(cards)
	return c.Status(fiber.StatusOK).JSON(serialization.HandToSer(hand))
}

func CalculateOdds(c *fiber.Ctx) error {
	deck := objects.NewDeck()
	requestBody := c.Context().Request.Body()
	var cardNames []string
	err := json.Unmarshal(requestBody, &cardNames)
	if err != nil {
		panic(err)
	}
	cards := make([]*objects.Card, 0)
	for _, cardName := range cardNames {
		cards = append(cards, deck.TakeName(cardName))
	}
	hand := funcs.EvaluateHand(cards)
	possibleHandsByRank := funcs.CalculateHandOdds(hand, cards, deck)
	result := fiber.Map{}
	for rank, possibleHands := range possibleHandsByRank {
		possibleHandsSers := make([]fiber.Map, len(possibleHands))
		for i, possibleHand := range possibleHands {
			possibleHandsSers[i] = serialization.PossibleHandToSer(possibleHand)
		}
		result[fmt.Sprint(rank)] = possibleHandsSers
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
