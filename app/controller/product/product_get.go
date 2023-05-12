package product

import (
	"synapsis-test/app/model"
	"synapsis-test/app/service"

	"github.com/gofiber/fiber/v2"
)

// GetProduct godoc
// @Summary      Get All product
// @Description  Mendapatkan semua list produk
// @Tags         Product
// @Accept       application/json
// @Produce		 application/json
// @Security ApiKeyAuth
// @Param page query int false "Page number start from zero"
// @Param size query int false "Size per page, default `0`"
// @Param sort query string false "Sort by field, adding dash (`-`) at the beginning means descending and vice versa"
// @Param fields query string false "Select specific fields with comma separated"
// @Param filters query string false "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)"
// @Success      200  {object}  []model.Product
// @Router       /product [get]
func GetProduct(c *fiber.Ctx) error {
	db := service.DB
	pg := service.PG

	mod := db.Model(&model.Product{})

	response := pg.With(mod).Request(c.Request()).Response(&[]model.Product{})

	return c.JSON(response)
}
