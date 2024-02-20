package router

import (
	"furyoftroy/pokerbuddy/v1/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type ApiRouter struct {
}

func (h ApiRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api", recover.New(), logger.New(), limiter.New())
	api.Post("/cards/evaluate", controllers.EvaluateCards)
	api.Post("/cards/calculateodds", controllers.CalculateOdds)
	api.Post("/cards/calculate", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from api",
		})
	})
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}
