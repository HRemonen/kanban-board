package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"

	jwtMiddleware "github.com/gofiber/jwt/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing or malformed JWT",
			"data":    nil,
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(401).JSON(fiber.Map{
		"status":  "error",
		"message": "Unauthorized action",
		"data":    nil,
	})
}
