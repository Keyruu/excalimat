package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/keyruu/excalimat-backend/config"
)

func parseBody(input interface{}, c *fiber.Ctx) error {
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on parsing body", "data": err})
	}
	return nil
}

func AdminCheck(c *fiber.Ctx) bool {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	groups := claims["groups"].([]interface{})
	for _, group := range groups {
		if group.(string) == config.ADMIN_GROUP {
			log.Println("is Admin")
			return true
		}
	}
	return false
}
