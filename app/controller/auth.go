package controller

import (
	"encoding/json"
	"fmt"
	"strings"

	"go-fiber/app/utils"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// LoginController handles user login requests.
// @Summary      User login
// @Description  Authenticates a user and returns a JWT token if credentials are valid.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      User  true  "User credentials"
// @Success      200   {object}  map[string]string  "JWT token"
// @Failure      400   {string}  string  "Invalid request body"
// @Failure      401   {string}  string  "Invalid credentials"
// @Failure      500   {string}  string  "No username found"
// @Router       /api/v1/login [post]
func LoginController(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", c.Body())

	var u User
	if err := json.Unmarshal(c.Body(), &u); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	fmt.Printf("The user request value %v", u)

	if u.Username == "admin" && u.Password == "123456" {
		tokenString, err := utils.CreateToken(u.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("No username found")
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenString})
	} else {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}
}

// ProtectedHandler is a handler for protected routes.
// It checks for a valid JWT token in the Authorization header.
// @Summary      Protected route
// @Description  This route is protected and requires a valid JWT token.
// @Tags         Protected
// @Accept       json
// @Produce      plain
// @Success      200  {string}  string  "Welcome to the protected area
// @Failure      401  {string}  string  "Unauthorized"
// @Security     ApiKeyAuth
// @Router       /api/v1/protected [get]
func ProtectedHandler(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing authorization header")
	}
	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid authorization header format")
	}
	tokenString := authHeader[len(bearerPrefix):]
	tokenString = strings.TrimSpace(tokenString)

	err := utils.VerifyToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}
	return c.SendString("Welcome to the protected area")
}
