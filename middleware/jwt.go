package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/keyruu/excalimat-backend/config"
	"github.com/keyruu/excalimat-backend/handler"
	"github.com/keyruu/excalimat-backend/sessions"
)

// Protected protect routes
func AuthRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		KeySetURL:    config.JwtKeyUrl,
		ErrorHandler: jwtError,
		Filter:       usesSession,
	})
}

func IsAdmin(c *fiber.Ctx) error {
	if !isAdmin(c) {
		return c.Status(fiber.StatusForbidden).JSON(handler.ErrorJSON("User is not in the Admin Group", nil))
	}
	return nil
}

func IsUser(c *fiber.Ctx) error {
	if !isUser(c) {
		return c.Status(fiber.StatusForbidden).JSON(handler.ErrorJSON("User is not in the User Group", nil))
	}
	return nil
}

func usesSession(c *fiber.Ctx) bool {
	sess, err := sessions.Store.Get(c)
	if err != nil {
		return false
	}
	return !sess.Fresh()
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(handler.ErrorJSON("Missing or malformed JWT", nil))
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(handler.ErrorJSON("Invalid or expired JWT", nil))
}

func isAdmin(c *fiber.Ctx) bool {
	return isGroup(config.AdminGroup, c)
}

func isUser(c *fiber.Ctx) bool {
	return isGroup(config.UserGroup, c)
}

func isGroup(group string, c *fiber.Ctx) bool {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	groups := claims["groups"].([]interface{})
	for _, adGroup := range groups {
		if adGroup.(string) == group {
			return true
		}
	}
	return false
}
