package router

import (
	"furyoftroy/pokerbuddy/v1/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type HttpRouter struct {
}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	group := app.Group("", recover.New(), logger.New(), cors.New(), csrf.New())
	group.Get("/", controllers.RenderIndex)
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
