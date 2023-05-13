package auth

import (
	"synapsis-test/app/middleware"
	"synapsis-test/app/model"
	"synapsis-test/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary      Login new costumer
// @Description  Login new costumer
// @Tags         Auth
// @Accept       application/json
// @Produce		 application/json
// @Param        data   body  auth.Payload  true  "Payload Login"
// @Success      200  {object}  model.Costumer
// @Router       /login [post]
func Login(c *fiber.Ctx) error {
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

// Payload struct
type Payload struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
