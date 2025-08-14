package routes

import (
	"go-fiber/app/controller"
	"go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	router := a.Group("/api/v1")

	// Routes for GET method:
	router.Get("/helloworld", middleware.JWTMiddleware, controller.HellowWorldController)
	router.Get("/users", middleware.JWTMiddleware, controller.UserController)

	// Routes for POST method:

}
