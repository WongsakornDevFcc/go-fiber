package controller

import (
	"encoding/json"
	"go-fiber/app/models"
	"go-fiber/pkg/utils"
	"go-fiber/platform/database"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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
//	@Param			user	body		models.SignIn				true	"User credentials"
//	@Success		200		{object}	map[string]string	"JWT token"
//	@Failure		400		{string}	string				"Invalid request body"
//	@Failure		401		{string}	string				"Invalid credentials"
//	@Failure		500		{string}	string				"No username found"
//	@Router			/api/v1/authentication/signin [post]
func LoginController(c *fiber.Ctx) error {
	var signIn = &models.SignIn{}

	if err := c.BodyParser(signIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	foundedUser, err := db.GetUserByEmail(signIn.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with the given email is not found",
		})
	}

	compareUserPassword := utils.ComparePasswords(foundedUser.PasswordHash, signIn.Password)
	if !compareUserPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "wrong user email address or password",
		})
	}

	// if signIn.Username == "admin" && signIn.Password == "123456" {
	tokenString, err := utils.CreateToken(signIn.Email, foundedUser.UserRole)
	refreshTokenString, err := utils.CreateRefreshToken(signIn.Email, foundedUser.UserRole)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("No username found")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"tokens": fiber.Map{
			"access":  tokenString,
			"refresh": refreshTokenString,
		},
		"user": fiber.Map{
			"username": signIn.Email,
			"role":     foundedUser.UserRole,
		}})
	// } else {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error":  true,
	// 		"msg":    "Authentication failed. Invalid or missing credentials.",
	// 		"status": fiber.StatusUnauthorized,
	// 	})
	// }
}

// ProtectedHandler is a handler for protected routes.
// It checks for a valid JWT token in the Authorization header.
//
//	@Summary		Protected route
//	@Description	This route is protected and requires a valid JWT token.
//	@Tags			Protected
//	@Accept			json
//	@Produce		json
//	@Params			token	body 								string	true	"JWT token"
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

// UserSignUp method to create a new user.
//
//	@Description	Create a new user.
//	@Summary		create a new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.SignUp	true	"Sign Up Body"
//	@Success		200		{object}	models.User
//	@Router			/api/v1/user/sign/up [post]
func UserSignUp(c *fiber.Ctx) error {
	signUp := &models.SignUp{}

	if err := c.BodyParser(signUp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(signUp); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking role from sign up data.
	role, err := utils.VerifyRole(signUp.UserRole)
	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new user struct.
	user := &models.User{}

	// Set initialized default data for user:
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.Email = signUp.Email
	user.PasswordHash = utils.GeneratePassword(signUp.Password)
	user.UserStatus = 1 // 0 == blocked, 1 == active
	user.UserRole = role

	// Validate user fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create a new user with validated data.
	if err := db.CreateUser(user); err != nil {
		// Return status 500 and create user process error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Delete password hash field from JSON view.
	user.PasswordHash = ""

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}
