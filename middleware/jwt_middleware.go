package middleware

import (
	"go-fiber/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing authorization header")
	}
	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid authorization header format")
	}

	tokenString := strings.TrimSpace(authHeader[len(bearerPrefix):])
	if err := utils.VerifyToken(tokenString); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "Invalid token",
			"status":  fiber.StatusUnauthorized,
		})
	}
	return c.Next()
}
