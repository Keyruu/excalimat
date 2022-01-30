package handler

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	hashedPin, err := bcrypt.GenerateFromPassword([]byte(RandomPIN()), bcrypt.DefaultCost)
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}

	input.PIN = string(hashedPin)

	result := database.DB.Create(&input)
	if result.Error != nil {
		log.Println(result.Error)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Created Account", "data": &input})
}
