package main

import (
	"go-fiber/app/database"
	"go-fiber/app/routes"
	_ "go-fiber/docs"
	"log"
	"os"

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

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database connection
	database.Connect()
	// database.Migrate()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: `*`,
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	app.Get("/swagger/*", swagger.HandlerDefault)

	routes.LoginRoute(v1)
	routes.ProtectedHandler(v1)
	routes.HelloWorld(v1)
	routes.Test(v1)
	routes.RefreshToken(v1)
	routes.UsersRoute(v1)

	port := os.Getenv("API_PORT")

	log.Println("Starting server on port:", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
