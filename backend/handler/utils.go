package handler

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/keyruu/excalimat-backend/config"
	"github.com/keyruu/excalimat-backend/database"
	"github.com/keyruu/excalimat-backend/model"
	"github.com/keyruu/excalimat-backend/sessions"
	"github.com/keyruu/excalimat-backend/storage"
	"github.com/keyruu/excalimat-backend/validation"
)

func parseBody(input interface{}, c *fiber.Ctx) error {
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	if err := validation.Validate.Struct(input); err != nil {
		log.Println("validation error")
		return err
	}
	return nil
}

func badRequest(err error, c *fiber.Ctx) error {
	return c.Status(400).JSON(ErrorJSON(err.Error(), nil))
}

func uploadFile(imageType string, c *fiber.Ctx) (string, error) {
	data, err := c.FormFile("document")
	if err != nil || data.Size == 0 {
		c.Status(fiber.StatusBadRequest).JSON(ErrorJSON("Couldn't get file from Form", err.Error()))
		return "", err
	}

	file, err := data.Open()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(ErrorJSON("Couldn't open file", err.Error()))
		return "", err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(ErrorJSON("Couldn't read file", err.Error()))
		return "", err
	}

	imagePath := "account/" + uuid.New().String() + filepath.Ext(data.Filename)

	err = storage.S3.Set(imagePath, bytes, time.Second)
	log.Println("Hello")
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(ErrorJSON("Couldn't upload file to S3", err.Error()))
		return "", err
	}

	err = file.Close()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(ErrorJSON("Couldn't close file", err.Error()))
		return "", err
	}

	return imagePath, nil
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

func IsAdmin(c *fiber.Ctx) bool {
	return isGroup(config.AdminGroup, c)
}

func IsUser(c *fiber.Ctx) bool {
	return isGroup(config.UserGroup, c)
}

func isGroup(group string, c *fiber.Ctx) bool {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	if claims["groups"] == nil {
		return false
	}

	groups := claims["groups"].([]interface{})
	for _, adGroup := range groups {
		if adGroup.(string) == group {
			return true
		}
	}
	return false
}
