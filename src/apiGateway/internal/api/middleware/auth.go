package middleware

import "github.com/gofiber/fiber/v2"

func LoggingMiddleware(c *fiber.Ctx) error {
	return c.Next()
}
func AuthMiddleware(c *fiber.Ctx) error {
	// Check for authentication token in the request
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Validate the token (this is just a placeholder, implement your own logic)
	if token != "valid-token" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden",
		})
	}

	return c.Next()
}
