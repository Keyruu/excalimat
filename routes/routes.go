package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/keyruu/excalimat-backend/handler"
	"github.com/keyruu/excalimat-backend/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api")
	v1 := api.Group("/v1", logger.New())
	v1.Get("/", middleware.AuthRequired(), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "success", "message": "Guten Abend. Dubinski.", "data": nil})
	})

	// Auth
	pin := v1.Group("/pin")
	pin.Post("/login", handler.Login)
	pin.Post("/set", middleware.AuthRequired(), handler.SetPIN)

	account := v1.Group("/account")
	account.Get("/", handler.GetAllAccounts)
	account.Get("/:id", handler.GetAccount)
	account.Post("/", handler.CreateAccount)
	account.Patch("/:id", middleware.AuthRequired(), middleware.AdminCheck, handler.UpdateAccount)
	account.Delete("/:id", middleware.AuthRequired(), middleware.AdminCheck, handler.DeleteAccount)
	account.Post("/signup", middleware.AuthRequired(), middleware.UserCheck, handler.SignUp)

	product := v1.Group("/product")
	product.Get("/", handler.GetAllProducts)
	product.Get("/:id", handler.GetProduct)
	product.Post("/", middleware.AuthRequired(), middleware.AdminCheck, handler.CreateProduct)
	product.Patch("/:id", middleware.AuthRequired(), middleware.AdminCheck, handler.UpdateProduct)
	product.Delete("/:id", middleware.AuthRequired(), middleware.AdminCheck, handler.DeleteProduct)

	purchase := v1.Group("/purchase")
	purchase.Get("/", middleware.AuthRequired(), middleware.AdminCheck, handler.GetAllPurchases)
	purchase.Get("/me", middleware.AuthRequired(), handler.GetMyPurchases)
	purchase.Get("/:id", middleware.AuthRequired(), middleware.AdminCheck, handler.GetPurchase)
	purchase.Post("/", middleware.AuthRequired(), handler.MakePurchase)
	purchase.Delete("/:id", middleware.AuthRequired(), handler.DeletePurchase)

	// // User
	// user := api.Group("/user")
	// user.Get("/:id", handler.GetUser)
	// user.Post("/", handler.CreateUser)
	// user.Patch("/:id", middleware.AuthRequired(), handler.UpdateUser)
	// user.Delete("/:id", middleware.AuthRequired(), handler.DeleteUser)

	// // Product
	// product := api.Group("/product")
	// product.Get("/", handler.GetAllProducts)
	// product.Get("/:id", handler.GetProduct)
	// product.Post("/", middleware.AuthRequired(), handler.CreateProduct)
	// product.Delete("/:id", middleware.AuthRequired(), handler.DeleteProduct)
}
