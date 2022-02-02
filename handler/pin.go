package handler

import (
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/keyruu/excalimat-backend/database"
	"github.com/keyruu/excalimat-backend/sessions"
	"golang.org/x/crypto/bcrypt"
)

const SessionUserId = "user_id"

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
	type PinInput struct {
		PIN string `json:"pin"`
	}

	db := database.DB

	account, err := CurrentAccount(c)
	if err != nil || account == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ErrorJSON("Couldn't get account from database", err))
	}

	var input PinInput

	err = parseBody(&input, c)
	if err != nil {
		return err
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
	type LoginInput struct {
		AccountID uint   `json:"account_id"`
		PIN       string `json:"pin"`
	}

	var input LoginInput

	err := parseBody(&input, c)
	if err != nil {
		return err
	}
	accountId := input.AccountID
	pin := input.PIN

	account, err := getAccountByID(accountId)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(ErrorJSON("Couldn't get account from database", err))
	}

	if !CheckHash(pin, account.PIN) {
		return c.Status(fiber.StatusUnauthorized).JSON(ErrorJSON("Invalid password", nil))
	}

	session, err := sessions.Store.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	defer session.Save()

	session.Set(SessionUserId, account.ID)

	return c.JSON(SuccessJSON("Login was successful", account))
}
