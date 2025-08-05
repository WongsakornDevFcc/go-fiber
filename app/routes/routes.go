package routes

import (
	"go-fiber/app/controller"

	"github.com/gofiber/fiber/v2"
)

func LoginRoute(router fiber.Router) {
	router.Get("/login", controller.LoginController)
}
func RegisterRoute(router fiber.Router) {
	router.Get("/healthcheck", controller.HealthCheckController)
}

func HelloWorld(router fiber.Router) {
	router.Get("/helloworld", controller.HellowWorldController)
}

func Test(router fiber.Router) {
	router.Get("/test", controller.TestController)
}
