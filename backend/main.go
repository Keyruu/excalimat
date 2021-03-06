package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet/v2"
	"github.com/keyruu/excalimat-backend/database"
	"github.com/keyruu/excalimat-backend/model"
	"github.com/keyruu/excalimat-backend/routes"
	"github.com/keyruu/excalimat-backend/storage"
	"github.com/keyruu/excalimat-backend/validation"
)

func main() {
	// Start a new fiber app
	app := fiber.New(fiber.Config{
		ReadBufferSize: 8192,
	})

	validation.InitValidation()

	storage.InitStorage()

	database.Connect()
	database.DB.AutoMigrate(&model.Account{}, &model.Product{}, &model.Purchase{})

	// Use middlewares for each route
	app.Use(
		helmet.New(), // add Helmet middleware
	)

	app.Use(cors.New())

	// app.Use(
	// 	csrf.New(), // add CSRF middleware
	// )

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Hier spricht Edgar Wallace sein Nachbar!")
		return err
	})

	app.Get("/health", HealthCheck)

	routes.SetupRoutes(app)

	// Listen on PORT 300
	app.Listen(":3000")
}

func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
