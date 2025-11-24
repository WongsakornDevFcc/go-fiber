package routes

import (
	"go-fiber/app/controller"
	"go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func PrivateRoutes(a *fiber.App, rdb *redis.Client) {
	router := a.Group("/api/v1",
		middleware.JWTProtected(),
		middleware.TokenBucketMiddleware(middleware.TokenBucketConfig{
			Rate:  1, // 1 token per second
			Burst: 3, // max 3 requests at once
			Redis: rdb,
		}),
	)

	// Routes for GET method:
	router.Get("/helloworld", middleware.JWTProtected(), controller.HellowWorldController)
	router.Get("/user", middleware.JWTProtected(), controller.GetUsers)

	// Routes for POST method:

}
