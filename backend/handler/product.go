package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keyruu/excalimat-backend/database"
	"github.com/keyruu/excalimat-backend/model"
	"github.com/keyruu/excalimat-backend/validation"
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

func UpdateProduct(c *fiber.Ctx) error {
	db := database.DB
	var oldProduct model.Product
	var product model.Product

	err := parseBody(&product, c)
	if err != nil {
		return badRequest(err, c)
	}

	result := db.Where("id = ?", c.Params("id")).Find(&oldProduct)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err := validation.Validate.Struct(oldProduct); err != nil {
		return c.Status(404).JSON(ErrorJSON("Product does not exist", nil))
	}

	oldProduct.Price = product.Price
	oldProduct.BundleSize = product.BundleSize
	oldProduct.Type = product.Type
	oldProduct.Name = product.Name
	oldProduct.Picture = product.Picture

	result = db.Save(&oldProduct)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(SuccessJSON("Updated product", oldProduct))
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var product model.Product
	db.First(&product, id)
	if err := validation.Validate.Struct(product); err != nil {
		return c.Status(404).JSON(ErrorJSON("No product found", nil))
	}
	db.Delete(&product)
	return c.JSON(SuccessJSON("Product successfully deleted", product))
}
