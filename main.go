package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keyruu/excalimat-backend/database"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	database.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is TEST!")
		return err
	})

	// Listen on PORT 300
	app.Listen(":3000")
}
