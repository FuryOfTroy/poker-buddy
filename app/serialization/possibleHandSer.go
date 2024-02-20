package serialization

import (
	"furyoftroy/pokerbuddy/v1/objects"

	"github.com/gofiber/fiber/v2"
)

func PossibleHandToSer(possibleHand *objects.PossibleHand) fiber.Map {
	outs := make([]fiber.Map, 0)
	for _, out := range possibleHand.GetOuts() {
		outs = append(outs, CardToSer(out))
	}
	return fiber.Map{"hand": HandToSer(possibleHand.GetHand()), "outs": outs, "print": possibleHand.Print()}
}
