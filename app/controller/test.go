package controller

import (
	"github.com/gofiber/fiber/v2"
)

// Test
//
//	@Summary		Show the test to server.
//	@Description	get test of server.
//	@Tags			Test
//	@Accept			*/*
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	"test"
//	@Router			/api/v1/test [get]
func TestController(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "test"})
}
