package serialization

import (
	"furyoftroy/pokerbuddy/v1/objects"

	"github.com/gofiber/fiber/v2"
)

func CardToSer(card *objects.Card) fiber.Map {
	return fiber.Map{"value": card.GetValue(), "suit": card.GetSuit(), "print": card.Print()}
}
