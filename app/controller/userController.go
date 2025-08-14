package controller

import (
	"github.com/gofiber/fiber/v2"
)

// User requests.
//
//	@Summary		Show the Users list to server.
//	@Description	get Users list of server.
//	@Tags			Users
//	@Accept			*/*
//	@Produce		json
//	@Success		200	"user list"
//	@Security		BearerAuth
//	@Router			/api/v1/users [get]
func UserController(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello",
		"content": "test",
	})

}
