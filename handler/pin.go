package handler

import (
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/keyruu/excalimat-backend/database"
	"golang.org/x/crypto/bcrypt"
)

type PinInput struct {
	AccountID uint   `json:"account_id"`
	PIN       string `json:"pin"`
}

// CheckPasswordHash compare password with hash
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

const letterBytes = "1234567890"

func RandomPIN() string {
	b := make([]byte, 4)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func SetPIN(c *fiber.Ctx) error {
	db := database.DB

	var input PinInput

	err := parseBody(&input, c)
	if err != nil {
		return err
	}

	account, err := getAccountByID(input.AccountID)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Couldn't get account from database", "data": err})
	}

	hashedPin, err := bcrypt.GenerateFromPassword([]byte(input.PIN), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	account.PIN = string(hashedPin)

	db.Save(&account)

	return c.SendStatus(fiber.StatusNoContent)
}

// Login get user and password
func Login(c *fiber.Ctx) error {
	var input PinInput

	err := parseBody(&input, c)
	if err != nil {
		return err
	}
	accountId := input.AccountID
	pin := input.PIN

	account, err := getAccountByID(accountId)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Couldn't get account from database", "data": err})
	}

	if !CheckHash(pin, account.PIN) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	session, err := sessionStore.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	defer session.Save()

	session.Set("user_id", account.ID)

	return c.JSON(fiber.Map{"status": "success", "message": "Login was successful", "data": account})
}
