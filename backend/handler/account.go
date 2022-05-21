package handler

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/keyruu/excalimat-backend/database"
	"github.com/keyruu/excalimat-backend/model"
	"github.com/keyruu/excalimat-backend/storage"
	"github.com/keyruu/excalimat-backend/validation"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func getAccountByID(id uint) (*model.Account, error) {
	db := database.DB
	var account model.Account
	if err := db.First(&account, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func CreateAccount(c *fiber.Ctx) error {
	var input model.Account

	err := parseBody(&input, c)
	if err != nil {
		return badRequest(err, c)
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

func UploadAccountImage(c *fiber.Ctx) error {
	db := database.DB
	currentAccount, err := CurrentAccount(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorJSON("Couldn't get current account", err.Error()))
	}

	if fmt.Sprint(currentAccount.ID) != c.Params("id") || !IsAdmin(c) {
		return c.SendStatus(fiber.StatusForbidden)
	}

	var account *model.Account
	if fmt.Sprint(currentAccount.ID) == c.Params("id") {
		account = currentAccount
	} else {
		if err := db.First(&account, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(ErrorJSON("Account was not found", err.Error()))
		}
	}

	// Get first file from form field "document":
	imagePath, err := uploadFile("account", c)
	if err != nil {
		return nil
	}

	if account.Picture != "" {
		err = storage.S3.Delete(account.Picture)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorJSON("Couldn't delete old image", err.Error()))
		}
	}

	account.Picture = imagePath

	err = db.Save(account).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorJSON("Couldn't save account", err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(SuccessJSON("Uploaded image", "/image/"+imagePath))
}

func GetAllAccounts(c *fiber.Ctx) error {
	db := database.DB
	var accounts []model.Account

	result := db.Find(&accounts)

	if result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(SuccessJSON("Found accounts", accounts))
}

func GetAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var account model.Account

	if err := db.First(&account, id).Error; err != nil {
		return c.Status(404).JSON(ErrorJSON("No user found with ID", err.Error()))
	}
	return c.JSON(SuccessJSON("Product found", account))
}

func UpdateAccount(c *fiber.Ctx) error {
	db := database.DB
	var oldAccount model.Account
	var account model.Account

	err := parseBody(&account, c)
	if err != nil {
		return badRequest(err, c)
	}

	if err := db.First(&oldAccount, c.Params("id")).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err := validation.Validate.Struct(oldAccount); err != nil {
		return c.Status(404).JSON(ErrorJSON("Account does not exist", oldAccount))
	}

	oldAccount.Email = account.Email
	oldAccount.ExtID = account.ExtID
	oldAccount.Balance = account.Balance
	oldAccount.Name = account.Name
	oldAccount.Picture = account.Picture

	if err := db.Save(&oldAccount).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(SuccessJSON("Updated account", oldAccount))
}

func DeleteAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var account model.Account

	db.First(&account, id)

	db.Delete(&account)
	return c.JSON(SuccessJSON("User successfully deleted", nil))
}

func SignUp(c *fiber.Ctx) error {
	var input PinInput

	err := parseBody(&input, c)
	if err != nil {
		return badRequest(err, c)
	}

	db := database.DB

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	extId := claims["oid"].(string)

	var exists bool
	db.Model(model.Account{}).
		Select("count(*) > 0").
		Where("ext_id = ?", extId).
		First(&exists)

	if exists {
		return c.Status(fiber.StatusConflict).JSON(ErrorJSON("This account already exists", nil))
	}

	email := claims["email"].(string)
	name := claims["name"].(string)
	hash, err := bcrypt.GenerateFromPassword([]byte(input.PIN), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	account := model.Account{ExtID: extId, Email: email, Name: name, Balance: 0, PIN: string(hash)}

	db.Create(&account)

	return c.Status(fiber.StatusCreated).JSON(SuccessJSON("Account created", &account))
}
