package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}

func main() {

	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
