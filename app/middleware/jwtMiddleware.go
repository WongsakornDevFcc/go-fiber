package middleware

import (
	"fmt"
	"go-fiber/app/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// JWTMiddleware checks for a valid JWT token in the Authorization header
func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	fmt.Println("Authorization header:", authHeader)
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing authorization header")
	}
	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid authorization header format")
	}

	tokenString := strings.TrimSpace(authHeader[len(bearerPrefix):])
	fmt.Println("Token string:", tokenString)
	if err := utils.VerifyToken(tokenString); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}
	return c.Next()
}
