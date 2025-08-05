package main

import (
	"go-fiber/app/routes"
	_ "go-fiber/docs"
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

//	@title			API by Fiber and Swagger
//	@version		1.0
//	@description	API management Server by Fiber | Doc by Swagger.

//	@contact.name	admin
//	@contact.url	http://test.com/support
//	@contact.email	admin@test.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// <--	@schemes	 https http -->

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
	routes.ProtectedHandler(v1)
	routes.HelloWorld(v1)
	routes.Test(v1)

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	// log.Println(os.Getenv("TEST_DATA"))

	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
