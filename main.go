package main

import (
	"log"

	"go-fiber/app/controller"
	"go-fiber/app/routes"
	_ "go-fiber/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//	@title			API by Fiber and Swagger
//	@version		1.0
//	@description	API management Server by Fiber | Doc by Swagger.

//	@contact.name	admin
//	@contact.url	http://test.com/support
//	@contact.email	admin@test.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@schemes	https http

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	app := fiber.New()

	app.Use(cors.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")
	app.Get("/swagger/*", swagger.HandlerDefault)

	routes.LoginRoute(v1)
	// routes.RegisterRoute(v1)
	routes.HelloWorld(v1)
	routes.Test(v1)

	// Unauthenticated route
	app.Get("/", controller.AccessibleController)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	// Restricted Routes
	// app.Get("/restricted", restricted)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
