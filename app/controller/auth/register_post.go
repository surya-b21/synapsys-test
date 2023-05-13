package auth

import (
	"net/mail"
	"synapsis-test/app/middleware"
	"synapsis-test/app/model"
	"synapsis-test/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary      Register new costumer
// @Description  Register new costumer
// @Tags         Auth
// @Accept       application/json
// @Produce		 application/json
// @Param        data   body  auth.Payload  true  "Payload Register"
// @Success      200  {object}  model.Costumer
// @Router       /register [post]
func Register(c *fiber.Ctx) error {
	var payload Payload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "invalid request",
		})
	}

	db := service.DB

	if validEmail := validEmailAddress(payload.Email); !validEmail {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "invalid email",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	stringPassword := string(hash)
	costumer := model.Costumer{
		CostumerAPI: model.CostumerAPI{
			Name:     &payload.Name,
			Email:    &payload.Email,
			Password: &stringPassword,
		},
	}

	db.Create(&costumer)

	token, err := middleware.Encode(&jwt.MapClaims{
		"email": payload.Email,
		"id":    costumer.ID.String(),
	}, 1000)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"token":      token,
		"expired_at": middleware.DefaultConfig.Expiry,
	})
}

func validEmailAddress(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	return true
}
