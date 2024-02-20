package app

import (
	"embed"
	"furyoftroy/pokerbuddy/v1/app/pkg/router"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/template/django/v3"
)

//go:embed static/*
var embedDirStatic embed.FS

func NewApplication() *fiber.App {
	engine := django.New("app/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "static",
	}))
	app.Get("/dashboard", monitor.New())
	router.InstallRouter(app)

	return app
}
