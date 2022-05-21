package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keyruu/excalimat-backend/database"
	"github.com/keyruu/excalimat-backend/model"
	"github.com/keyruu/excalimat-backend/storage"
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
	if err := db.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(ErrorJSON("No product found with ID", err))
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

	if err := db.First(&oldProduct, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorJSON("Couldn't find product", err.Error()))
	}

	if err := validation.Validate.Struct(oldProduct); err != nil {
		return c.Status(404).JSON(ErrorJSON("Product does not exist", err.Error()))
	}

	oldProduct.Price = product.Price
	oldProduct.BundleSize = product.BundleSize
	oldProduct.Type = product.Type
	oldProduct.Name = product.Name
	oldProduct.Picture = product.Picture

	if err := db.Save(&oldProduct).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(SuccessJSON("Updated product", oldProduct))
}

func UploadProductImage(c *fiber.Ctx) error {
	db := database.DB
	var product *model.Product

	if err := db.First(&product, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorJSON("Couldn't find product", err.Error()))
	}

	// Get first file from form field "document":
	imagePath, err := uploadFile("product", c)
	if err != nil {
		return nil
	}

	if product.Picture != "" {
		err = storage.S3.Delete(product.Picture)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorJSON("Couldn't delete old image", err.Error()))
		}
	}

	product.Picture = imagePath

	err = db.Save(product).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorJSON("Couldn't save account", err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(SuccessJSON("Uploaded image", "/image/"+imagePath))
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
