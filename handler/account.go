package handler

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/keyruu/excalimat-backend/database"
	"github.com/keyruu/excalimat-backend/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func getAccountByID(id uint) (*model.Account, error) {
	db := database.DB
	var account model.Account
	if err := db.Where("id = ?", id).Find(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func CreateAccount(c *fiber.Ctx) error {
	var input model.Account

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorJSON("Error on login request", err))
	}

	hashedPin, err := bcrypt.GenerateFromPassword([]byte(RandomPIN()), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	input.PIN = string(hashedPin)

	result := database.DB.Create(&input)
	if result.Error != nil {
		log.Println(result.Error)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(SuccessJSON("Created Account", &input))
}

func SignUp(c *fiber.Ctx) error {
	db := database.DB

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	extId := claims["oid"].(string)
	email := claims["email"].(string)
	name := claims["name"].(string)
	pin := RandomPIN()
	hash, err := bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	log.Println("The PIN is ", pin)

	account := model.Account{ExtID: extId, Email: email, Name: name, Balance: 0, PIN: string(hash)}

	db.Create(&account)

	return c.Status(fiber.StatusCreated).JSON(SuccessJSON("Account created", &account))
}
