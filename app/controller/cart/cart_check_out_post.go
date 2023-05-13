package cart

import (
	"synapsis-test/app/model"
	"synapsis-test/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// PostCartCheckOut godoc
// @Summary      Post checkout cart
// @Description  Checkout product in cart
// @Tags         Cart
// @Accept       application/json
// @Security 	ApiKeyAuth
// @Success      200  {object}  []model.Cart
// @Router       /cart/check-out [post]
func PostCartCheckOut(c *fiber.Ctx) error {
	db := service.DB

	// get costumer id
	claimData := c.Locals("jwtClaims")
	user := claimData.(jwt.MapClaims)

	cart := []model.Cart{}
	if rows := db.Where("costumer_id = ? AND is_checkout = ?", user["id"], false).Find(&cart).RowsAffected; rows < 1 {
		return c.JSON(fiber.Map{
			"message": "no product in cart",
		})
	}

	checkedOut := true
	for i := range cart {
		cart[i].IsCheckout = &checkedOut

		db.Model(&model.Cart{}).Where("id = ?", cart[i].ID).Update("is_checkout", true)
	}

	// db.Where("costumer_id = ? AND is_checkout = '0' AND deleted_at IS NULL", user["id"]).Updates(&cart)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "your product already checkout",
		"additional": fiber.Map{
			"invoice": cart[0].ID.String(),
		},
	})
}
