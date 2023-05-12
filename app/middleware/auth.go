package middleware

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Config for middleware
type Config struct {
	Filter       func(c *fiber.Ctx) bool
	Unauthorized fiber.Handler
	Decode       func(c *fiber.Ctx) (*jwt.MapClaims, error)
	Secret       string
	Expiry       int64
}

// DefaultConfig middleware
var DefaultConfig = Config{
	Filter:       nil,
	Decode:       nil,
	Unauthorized: nil,
	Secret:       "53cr3t!",
	Expiry:       60,
}

// configDefault for generating default config
func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return DefaultConfig
	}

	cfg := config[0]

	if cfg.Filter == nil {
		cfg.Filter = DefaultConfig.Filter
	}

	if cfg.Secret == "" {
		cfg.Secret = DefaultConfig.Secret
	}

	if cfg.Expiry == 0 {
		cfg.Expiry = DefaultConfig.Expiry
	}

	if cfg.Decode == nil {
		cfg.Decode = func(c *fiber.Ctx) (*jwt.MapClaims, error) {
			authHeader := c.Get("Authorization")

			if authHeader == "" {
				return nil, errors.New("Authorization header is required")
			}

			token, err := jwt.Parse(
				authHeader[7:],
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
					}

					return []byte(cfg.Secret), nil
				},
			)

			if err != nil {
				return nil, errors.New("Error parsing token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !(ok && token.Valid) {
				return nil, errors.New("Invalid token")
			}

			if expiresAt, ok := claims["exp"]; ok && int64(expiresAt.(float64)) < time.Now().UTC().Unix() {
				return nil, errors.New("token is expired")
			}

			return &claims, nil
		}
	}

	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	return cfg
}

// New to initialize middleware
func New(config Config) fiber.Handler {
	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {
		if cfg.Filter != nil && cfg.Filter(c) {
			return c.Next()
		}

		claims, err := cfg.Decode(c)

		if err == nil {
			c.Locals("jwtClaims", *claims)
			return c.Next()
		}

		return cfg.Unauthorized(c)
	}
}

// Encode to generating new token
func Encode(claims *jwt.MapClaims, expiryAfter int64) (string, error) {
	if expiryAfter == 0 {
		expiryAfter = DefaultConfig.Expiry
	}

	(*claims)["exp"] = time.Now().UTC().Unix() + expiryAfter

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(DefaultConfig.Secret))

	if err != nil {
		return "", errors.New("Error generating new token")
	}

	return signedToken, nil
}
