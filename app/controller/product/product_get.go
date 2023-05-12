package product

import (
	"synapsis-test/app/model"
	"synapsis-test/app/service"

	"github.com/gofiber/fiber/v2"
)

// GetProduct to get all product
func GetProduct(c *fiber.Ctx) error {
	db := service.DB
	pg := service.PG

	mod := db.Model(&model.Product{})

	response := pg.With(mod).Request(c.Request()).Response(&[]model.Product{})

	return c.JSON(response)
}
