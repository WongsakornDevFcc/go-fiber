package controller

import (
	"github.com/gofiber/fiber/v2"
)

// Helloworld
//
//	@Summary		Show the HelloWorld to server.
//	@Description	get Hello World of server.
//	@Tags			Test
//	@Accept			*/*
//	@Produce		json
//	@Success		200	"Hello world!!"
//	@Security		BearerAuth
//	@Router			/api/v1/helloworld [get]
func HellowWorldController(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello",
		"content": "test",
		"status":  fiber.StatusOK,
	})

}
