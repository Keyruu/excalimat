package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keyruu/excalimat-backend/database"
	"github.com/keyruu/excalimat-backend/model"
)

func GetAllProducts(c *fiber.Ctx) error {
	db := database.DB
	var products []model.Product

	result := db.Find(&products)

	if result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(SuccessJSON("Found products", products))
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var product model.Product
	db.Find(&product, id)
	if product.Name == "" {
		return c.Status(404).JSON(ErrorJSON("No product found with ID", nil))

	}
	return c.JSON(SuccessJSON("Product found", product))
}

func CreateProduct(c *fiber.Ctx) error {
	db := database.DB
	var product model.Product

	err := parseBody(&product, c)
	if err != nil {
		return err
	}

	result := db.Save(&product)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(SuccessJSON("Created product", product))
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var product model.Product
	db.First(&product, id)
	if product.Name == "" {
		return c.Status(404).JSON(ErrorJSON("No product found with ID", nil))

	}
	db.Delete(&product)
	return c.JSON(SuccessJSON("Product successfully deleted", nil))
}
