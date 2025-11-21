package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Test
//
//	@Summary		Show the test to server test.
//	@Description	get test of server.
//	@Tags			Test
//	@Accept			*/*
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	"test"
//	@Router			/api/v1/test [get]
func TestController(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Test successful",
		"content": "This is a test endpoint",
		"test":    "This is a test",
		"status":  fiber.StatusOK,
	})
}

// TestFastResponse represents the response structure for TestFastController
type TestFastResponse struct {
	Names []string `json:"names" example:"testfast1,testfast2,..."`
	Count int      `json:"count" example:"100"`
}

// TestFastController
//
//	@Summary		Generate list of testfast names.
//	@Description	Generate list of names from testfast1 to testfast100.
//	@Tags			Test
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	controller.TestFastResponse
//	@Router			/api/v1/testfast [get]
func TestFastController(c *fiber.Ctx) error {
	names := make([]string, 100)
	for i := 0; i < 100; i++ {
		names[i] = fmt.Sprintf("testfast%d", i+1)
	}

	return c.JSON(fiber.Map{
		"names": names,
		"count": 100,
	})
}
