package middleware

import "github.com/gofiber/fiber/v2"

// ClaimJWT function to claim jwt
func ClaimJWT(c *fiber.Ctx) error {
	claimData := c.Locals("jwtClaims")

	if claimData == nil {
		return c.JSON(fiber.ErrUnauthorized)
	}

	return c.Next()
}
