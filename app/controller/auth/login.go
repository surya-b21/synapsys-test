package auth

import (
	"synapsis-test/app/middleware"
	"synapsis-test/app/model"
	"synapsis-test/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Login function
func Login(c *fiber.Ctx) error {
	var payload LoginPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "invalid request",
		})
	}

	// hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	panic(err)
	// }

	db := service.DB

	costumer := model.Costumer{}
	if rows := db.First(&costumer, "email = ?", payload.Email).RowsAffected; rows < 1 {
		return c.Status(401).JSON(fiber.Map{
			"status":  "Unauthorized",
			"message": "Costumer not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*costumer.Password), []byte(payload.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"status":  "Unauthorized",
			"message": "Wrong password",
		})
	}

	token, err := middleware.Encode(&jwt.MapClaims{
		"email": payload.Email,
		"id":    costumer.ID.String(),
	}, middleware.DefaultConfig.Expiry)

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

// LoginPayload struct
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
