package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// HealthCheck
//
//	@Summary		Show the status of server.
//	@Description	get the status of server.
//	@Tags			CheckServer
//	@Accept			*/*
//	@Produce		json
//	@Success		200	"OK"
//	@Router			/api/v1/healthcheck [get]
func HealthCheckController(c *fiber.Ctx) error {
	msg := "OK"
	c.Status(http.StatusOK)
	return c.SendString(msg)
}
