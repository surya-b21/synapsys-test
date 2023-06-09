package router

import (
	authController "synapsis-test/app/controller/auth"
	"synapsis-test/app/controller/cart"
	"synapsis-test/app/controller/product"
	"synapsis-test/app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Handle app route
func Handle(app *fiber.App) {
	app.Use(cors.New())
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "This is test app",
		})
	})

	app.Post("/api/login", authController.Login)
	app.Post("/api/register", authController.Register)
	auth := middleware.New(middleware.DefaultConfig)

	api := app.Group("/api", auth, middleware.ClaimJWT)
	api.Get("/product", product.GetProduct)

	api.Get("/cart", cart.GetCart)
	api.Post("/cart", cart.PostCart)
	api.Post("/cart/check-out", cart.PostCartCheckOut)
	api.Delete("/cart/:id", cart.DeleteCart)
}
