package cart

import (
	"synapsis-test/app/model"
	"synapsis-test/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// DeleteCart godoc
// @Summary      Delete cart
// @Description  Delete product to cart
// @Tags         Cart
// @Accept       application/json
// @Security 	ApiKeyAuth
// @Param        id path string true "Cart ID"
// @Success      200  {object}  []model.CartAPI
// @Router       /cart/{id} [delete]
func DeleteCart(c *fiber.Ctx) error {
	db := service.DB

	_, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "invalid",
			"message": "invalid cart id",
		})
	}

	db.Where("id = ?", c.Params("id")).Delete(&model.Cart{})

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "product has been deleted from cart",
	})
}
