package controller

import (
	"go-fiber/platform/database"

	"github.com/gofiber/fiber/v2"
)

// User requests.
//
//	@Summary		Show the Users list to server.
//	@Description	get Users list of server.
//	@Tags			Users
//	@Accept			*/*
//	@Produce		json
//	@Success		200	"user list"
//	@Security		BearerAuth
//	@Router			/api/v1/users [get]
func GetUsers(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all users.
	users, err := db.GetUsers()
	if err != nil {
		// Return, if users not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user were not found",
			"count": 0,
			"users": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(users),
		"users": users,
	})

}
