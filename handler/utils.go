package handler

import "github.com/gofiber/fiber/v2"

func parseBody(input interface{}, c *fiber.Ctx) error {
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on parsing body", "data": err})
	}
	return nil
}
