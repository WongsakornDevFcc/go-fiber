package routes

import (
	"go-fiber/app/controller"
	"go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App, rdb *redis.Client) {
	// Create routes group.
	router := a.Group("/api/v1",
		middleware.TokenBucketMiddleware(middleware.TokenBucketConfig{
			Rate:  1, // 1 token per second
			Burst: 3, // max 3 requests at once
			Redis: rdb,
		}))

	// Routes for GET method:
	router.Get("/test", controller.TestController)
	router.Get("/protected", controller.ProtectedHandler)

	// Routes for POST method:
	router.Post("/authentication/signin", controller.LoginController)
	router.Post("/authentication/refresh", controller.RefreshTokenController)
	router.Post("/user/sign/up", controller.UserSignUp)

}
