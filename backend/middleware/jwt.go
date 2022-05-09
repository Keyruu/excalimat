package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
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

func AdminCheck(c *fiber.Ctx) error {
	if !handler.IsAdmin(c) {
		return c.Status(fiber.StatusForbidden).JSON(handler.ErrorJSON("User is not in the Admin Group", nil))
	}
	return c.Next()
}

func UserCheck(c *fiber.Ctx) error {
	if !handler.IsUser(c) {
		return c.Status(fiber.StatusForbidden).JSON(handler.ErrorJSON("User is not in the User Group", nil))
	}
	return c.Next()
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
