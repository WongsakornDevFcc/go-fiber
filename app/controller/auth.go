package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"
)

// LoginController handles user authentication.
//
// @Summary      Auth to server.
// @Description  Authorization of server.
// @Tags         Authentication
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        user  formData  string  true  "Username"
// @Param        pass  formData  string  true  "Password"
// @Success      200   {object}  map[string]string  "token"
// @Failure      401   {string}  string  "Unauthorized"
// @Failure      500   {string}  string  "Internal Server Error"
// @Router       /api/v1/login [post]
func LoginController(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	if user != "admin" || pass != "Vcxz1234!" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := jwt.MapClaims{
		"name":  "admin",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

// AccessibleController returns a public message.
//
// @Summary      Public access
// @Description  This endpoint is accessible without authentication.
// @Tags         Public
// @Accept       json
// @Produce      plain
// @Success      200  {string}  string  "Accessible"
// @Router       /api/v1/accessible [get]
func AccessibleController(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

// RestrictedController returns a message for authenticated users.
//
// @Summary      Restricted access
// @Description  This endpoint requires a valid JWT token.
// @Tags         Protected
// @Accept       json
// @Produce      plain
// @Success      200  {string}  string  "Welcome <name>"
// @Failure      401  {string}  string  "Unauthorized"
// @Security     ApiKeyAuth
// @Router       /api/v1/restricted [get]
func RestrictedController(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
