package cart

import (
	"synapsis-test/app/model"
	"synapsis-test/app/service"

	"github.com/gofiber/fiber/v2"
)

// PostCart godoc
// @Summary      Post cart
// @Description  Add product to cart
// @Tags         Cart
// @Accept       application/json
// @Produce		 application/json
// @Security 	ApiKeyAuth
// @Param        data   body  []model.CartAPI  true  "Payload Cart"
// @Success      200  {object}  []model.CartAPI
// @Router       /cart [post]
func PostCart(c *fiber.Ctx) error {
	payload := []model.CartAPI{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "invalid request",
		})
	}

	db := service.DB
	notCheckout := false
	cart := []model.Cart{}
	for _, v := range payload {
		cart = append(cart, model.Cart{
			CartAPI: model.CartAPI{
				ProductID:  v.ProductID,
				CostumerID: v.CostumerID,
				IsCheckout: &notCheckout,
			},
		})
	}
	if err := db.CreateInBatches(&cart, 1000).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "failed",
			"message": "failed to add product to cart",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "success add product to cart",
	})
}
