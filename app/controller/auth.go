package controller

import (
	"encoding/json"
	"go-fiber/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type TokenRequest struct {
	Refresh string `json:"refresh"`
}

// LoginController handles user login requests.
//
//	@Summary		User login
//	@Description	Authenticates a user and returns a JWT token if credentials are valid.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		User				true	"User credentials"
//	@Success		200		{object}	map[string]string	"JWT token"
//	@Failure		400		{string}	string				"Invalid request body"
//	@Failure		401		{string}	string				"Invalid credentials"
//	@Failure		500		{string}	string				"No username found"
//	@Router			/api/v1/authentication/signin [post]
func LoginController(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	var u User
	if err := json.Unmarshal(c.Body(), &u); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if u.Username == "admin" && u.Password == "123456" {
		tokenString, err := utils.CreateToken(u.Username, "admin")
		refreshTokenString, err := utils.CreateRefreshToken(u.Username, "admin")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("No username found")
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"tokens":  tokenString,
			"refresh": refreshTokenString,
			"user": fiber.Map{
				"username": u.Username,
				"role":     "admin",
			}})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "Authentication failed. Invalid or missing credentials.",
			"status":  fiber.StatusUnauthorized,
		})
	}
}

// ProtectedHandler is a handler for protected routes.
// It checks for a valid JWT token in the Authorization header.
//
//	@Summary		Protected route
//	@Description	This route is protected and requires a valid JWT token.
//	@Tags			Protected
//	@Accept			json
//	@Produce		json
//	@Params			token	body 				string	true	"JWT token"
//	@Success		200		{string}	string	"Welcome to the protected area
//	@Failure		401		{string}	string	"Unauthorized"
//	@Security		ApiKeyAuth
//	@Router			/api/v1/protected [get]
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

// RefreshToken handles token refresh requests.
//
//	@Summary		Token refresh
//	@Description	Refreshes a JWT token if the provided token is valid.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			token	body		TokenRequest		true	"JWT token"
//	@Success		200		{object}	map[string]string	"New JWT token"
//	@Failure		400		{string}	string				"Invalid request body"
//	@Failure		401		{string}	string				"Invalid token"
//	@Failure		500		{string}	string				"Failed to refresh token"
//	@Router			/api/v1/authentication/refresh [post]
func RefreshTokenController(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	var u TokenRequest
	if err := json.Unmarshal(c.Body(), &u); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	username, role, err := utils.VerifyRefreshToken(u.Refresh)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid refresh token")
	}

	newToken, err := utils.CreateToken(username, role)
	newRefreshToken := u.Refresh

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create new token")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"tokens":  newToken,
		"refresh": newRefreshToken,
		"user": fiber.Map{
			"username": username,
			"role":     role,
		},
	})
}
