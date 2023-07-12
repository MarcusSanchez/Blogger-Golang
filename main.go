package main

import (
	"blogger/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"log"
)

func main() {
	app := appInit()
	port := "3000"
	router.Router(app)
	log.Fatal(app.Listen(":" + port))
}

func appInit() *fiber.App {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
		return c.Next()
	})
	return app
}
