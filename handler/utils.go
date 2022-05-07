package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/keyruu/excalimat-backend/database"
	"github.com/keyruu/excalimat-backend/model"
	"github.com/keyruu/excalimat-backend/sessions"
)

func parseBody(input interface{}, c *fiber.Ctx) error {
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	if _, err := govalidator.ValidateStruct(input); err != nil {
		return err
	}
	return nil
}

func badRequest(err error, c *fiber.Ctx) error {
	return c.Status(400).JSON(ErrorJSON(err.Error(), nil))
}

func CurrentAccount(c *fiber.Ctx) (*model.Account, error) {
	db := database.DB
	var account *model.Account
	user := c.Locals("user")

	if user != nil {
		token := user.(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		oid := claims["oid"].(string)

		result := db.Where("ext_id = ?", oid).First(&account)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	sess, err := sessions.Store.Get(c)
	if err != nil {
		return nil, err
	}

	if !sess.Fresh() {
		userId := sess.Get(SessionUserId).(uint)

		account, err = getAccountByID(userId)
		if err != nil {
			return nil, err
		}
	}

	return account, nil
}

func SuccessJSON(message string, data interface{}) fiber.Map {
	return fiber.Map{"status": "success", "message": message, "data": data}
}

func ErrorJSON(message string, data interface{}) fiber.Map {
	return fiber.Map{"status": "error", "message": message, "data": data}
}
