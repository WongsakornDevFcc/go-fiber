package routes

import (
	"go-fiber/app/controller"
	"go-fiber/app/middleware"

	"github.com/gofiber/fiber/v2"
)

// public routes
func Test(router fiber.Router) {
	router.Get("/test", controller.TestController)
}

func LoginRoute(router fiber.Router) {
	router.Post("/authentication/signin", controller.LoginController)
}
func ProtectedHandler(router fiber.Router) {
	router.Get("/protected", controller.ProtectedHandler)
}

// private routes
func HelloWorld(router fiber.Router) {
	router.Get("/helloworld", middleware.JWTMiddleware, controller.HellowWorldController)
}
