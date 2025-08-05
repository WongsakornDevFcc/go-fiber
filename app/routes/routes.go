package routes

import (
	"go-fiber/app/controller"
	"go-fiber/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func LoginRoute(router fiber.Router) {
	router.Post("/login", controller.LoginController)
}
func ProtectedHandler(router fiber.Router) {
	router.Get("/protected", controller.ProtectedHandler)
}

func HelloWorld(router fiber.Router) {
	router.Get("/helloworld", middleware.JWTMiddleware, controller.HellowWorldController)
}

func Test(router fiber.Router) {
	router.Get("/test", middleware.JWTMiddleware, controller.TestController)
}
