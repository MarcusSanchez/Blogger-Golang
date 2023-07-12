package router

import (
	"blogger/handlers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/", handlers.GetRoot)
	app.Get("/create", handlers.GetCreate)
	app.Post("/create", handlers.PostCreate)
}
