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
//	@Router			/api/v1/helloworld [get]
func HellowWorldController(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hello world!!"})

}
