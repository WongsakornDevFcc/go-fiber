package routes

import (
	"go-fiber/app/controller"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	router := a.Group("/api/v1")

	// Routes for GET method:
	router.Get("/test", controller.TestController)
	router.Get("/protected", controller.ProtectedHandler)

	// Routes for POST method:
	router.Post("/authentication/signin", controller.LoginController)
	router.Post("/authentication/refresh", controller.RefreshTokenController)

}
