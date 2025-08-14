package main

import (
	_ "go-fiber/docs"
	"go-fiber/middleware"
	"go-fiber/pkg/configs"
	"go-fiber/pkg/routes"
	"go-fiber/pkg/utils"

	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
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

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	utils.StartServer(app)

}
