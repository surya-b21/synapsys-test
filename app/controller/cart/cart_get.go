package cart

import (
	"synapsis-test/app/model"
	"synapsis-test/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// GetCart godoc
// @Summary      Get all cart
// @Description  Get all product in cart
// @Tags         Cart
// @Accept       application/json
// @Produce		 application/json
// @Security ApiKeyAuth
// @Param page query int false "Page number start from zero"
// @Param size query int false "Size per page, default `0`"
// @Param sort query string false "Sort by field, adding dash (`-`) at the beginning means descending and vice versa"
// @Param fields query string false "Select specific fields with comma separated"
// @Param filters query string false "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)"
// @Success      200  {object}  []model.Cart
// @Router       /cart [get]
func GetCart(c *fiber.Ctx) error {
	db := service.DB
	pg := service.PG

	// get costumer id
	claimData := c.Locals("jwtClaims")
	user := claimData.(jwt.MapClaims)

	mod := db.Model(&model.Cart{}).Joins("Product").Joins("Costumer").Where("costumer_id = ?", user["id"])

	response := pg.With(mod).Request(c.Request()).Response(&[]model.Cart{})

	return c.JSON(response)
}
