package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/keyruu/excalimat-backend/storage"
)

func GetImage(c *fiber.Ctx) error {
	imageType := c.Params("type")
	id := c.Params("id")

	log.Println(id)

	if imageType == "" || id == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	file, err := storage.S3.Get(imageType + "/" + id)
	if err != nil || len(file) == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	c.Set("Content-Type", "application/octet-stream")
	return c.Status(fiber.StatusOK).Send(file)
}
