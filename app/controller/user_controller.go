package controller

import (
	"go-fiber/platform/database"
	"math"

	"github.com/gofiber/fiber/v2"
)

// GetUsers returns a paginated list of users.
//
//	@Summary		Show list of users
//	@Description	Get a paginated list of users from the server
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int	false	"Page number (default 1)"
//	@Param			limit	query		int	false	"Page size (default 10)"
//	@Security		BearerAuth
//	@Router			/api/v1/user [get]
func GetUsers(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 25)
	offset := (page - 1) * limit

	if page < 1 || limit < 1 {
		return c.Status(fiber.StatusBadRequest).JSON((fiber.Map{
			"error":   true,
			"message": "page or limit must have value",
			"count":   0,
			"user":    nil,
		}))
	}

	// Get all users.
	users, total, err := db.GetUsers(limit, offset)
	if err != nil {
		// Return, if users not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "user were not found",
			"count":   0,
			"users":   nil,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"page":    page,
		"limit":   limit,
		// "count":      len(users),
		"total":      total,
		"totalPages": totalPages,
		"users":      users,
	})
}
