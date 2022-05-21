package handler

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/keyruu/excalimat-backend/database"
	"github.com/keyruu/excalimat-backend/model"
	"gorm.io/gorm"
)

func MakePurchase(c *fiber.Ctx) error {
	type ProductInput struct {
		ProductIds []uint `json:"productIds" validate:"required,min=1"`
	}

	var input ProductInput

	err := parseBody(&input, c)
	if err != nil {
		return badRequest(err, c)
	}

	log.Println(input)

	db := database.DB

	account, err := CurrentAccount(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var purchases []model.Purchase
	err = db.Transaction(func(tx *gorm.DB) error {
		for _, id := range input.ProductIds {
			var product model.Product

			if err := tx.First(&product, id).Error; err != nil {
				return err
			}

			purchase := model.Purchase{Account: *account, Product: product, PaidPrice: product.Price}

			if err := tx.Create(&purchase).Error; err != nil {
				// return any error will rollback
				return err
			}

			account.Balance -= product.Price
			if err := tx.Save(&account).Error; err != nil {
				return err
			}

			purchases = append(purchases, purchase)
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(SuccessJSON("Made purchase", purchases))
}

func GetAllPurchases(c *fiber.Ctx) error {
	db := database.DB
	var purchases []model.Purchase

	result := db.Preload("Account").Preload("Product").Find(&purchases)

	if result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(SuccessJSON("Found purchases", purchases))
}

func GetPurchase(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var purchase model.Purchase

	if err := db.Preload("Account").Preload("Product").First(&purchase, id).Error; err != nil {
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
	var purchases []model.Purchase

	result := db.Where("account_id = ?", account.ID).Preload("Product").Find(&purchases)

	if result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(SuccessJSON("Found purchases", purchases))
}

func DeletePurchase(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var purchase model.Purchase

	if err := db.First(&purchase, id).Error; err != nil {
		return c.Status(404).JSON(ErrorJSON("No product found", err))
	}

	threshhold := purchase.CreatedAt.Add(time.Duration(5) * time.Minute)
	if time.Now().Before(threshhold) || IsAdmin(c) {
		err := db.Transaction(func(tx *gorm.DB) error {
			purchase.Account.Balance += purchase.PaidPrice

			if err := tx.Save(&purchase.Account).Error; err != nil {
				return err
			}
			if err := tx.Delete(&purchase).Error; err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	} else {
		return c.Status(fiber.StatusForbidden).JSON(ErrorJSON("Not allowed to refund after 5 minutes", nil))
	}

	return c.JSON(SuccessJSON("Purchase successfully deleted", purchase))
}
