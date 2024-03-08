package serialization

import (
	"furyoftroy/pokerbuddy/v1/objects"

	"github.com/gofiber/fiber/v2"
)

func HandToSer(hand *objects.Hand) fiber.Map {
	cards := make([]fiber.Map, 0)
	for _, card := range hand.GetCards() {
		cards = append(cards, CardToSer(card))
	}
	return fiber.Map{"rank": hand.GetRank(), "cards": cards, "print": hand.String()}
}
