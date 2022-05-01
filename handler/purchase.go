package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/keyruu/excalimat-backend/database"
	"github.com/keyruu/excalimat-backend/model"
	"gorm.io/gorm"
)

func MakePurchase(c *fiber.Ctx) error {
	type PurchaseInput struct {
		products []model.Product
	}

	var purchases PurchaseInput

	err := parseBody(&purchases, c)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	db := database.DB

	account, err := CurrentAccount(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		for _, product := range purchases.products {
			purchase := model.Purchase{AccountID: account.ID, ProductID: product.ID, PaidPrice: product.Price}

			if err := tx.Create(&purchase).Error; err != nil {
				// return any error will rollback
				return err
			}

			account.Balance -= product.Price
			if err := tx.Save(&account).Error; err != nil {
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func GetAllPurchases(c *fiber.Ctx) error {
	db := database.DB
	var purchases []model.Purchase

	result := db.Find(&purchases)

	if result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(SuccessJSON("Found purchases", purchases))
}

func GetPurchase(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var purchase model.Purchase

	result := db.Find(&purchase, id)

	if result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(SuccessJSON("Found purchase", purchase))
}

func GetMyPurchases(c *fiber.Ctx) error {
	account, err := CurrentAccount(c)
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}

	db := database.DB
	var purchases []model.Account

	result := db.Where("account_id = ?", account.ID).Find(&purchases)

	if result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(SuccessJSON("Found purchases", purchases))
}

func RefundPurchase(c *fiber.Ctx) error {
	account, err := CurrentAccount(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	log.Println(account.ID)

	return nil
}
