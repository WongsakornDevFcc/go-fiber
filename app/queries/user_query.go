package queries

import (
	"go-fiber/app/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// UserQueries struct for queries from User model.
type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) GetUsers(limit, offset int) ([]models.User, int, error) {
	users := []models.User{}

	query := `SELECT * FROM users ORDER BY created_at LIMIT $1 OFFSET $2`
	countQuery := `SELECT COUNT(*) FROM users`

	var total int
	if err := q.Get(&total, countQuery); err != nil {
		return nil, 0, err
	}

	if err := q.Select(&users, query, limit, offset); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// GetUserByID query for getting one User by given ID.
func (q *UserQueries) GetUserByID(id uuid.UUID) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Define query string.
	query := `SELECT * FROM users WHERE id = $1`

	// Send query to database.
	err := q.Get(&user, query, id)
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// GetUserByEmail query for getting one User by given Email.
func (q *UserQueries) GetUserByEmail(email string) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Define query string.
	query := `SELECT * FROM users WHERE email = $1`

	// Send query to database.
	err := q.Get(&user, query, email)
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// CreateUser query for creating a new user by given email and password hash.
func (q *UserQueries) CreateUser(u *models.User) error {
	// Define query string.
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database.
	_, err := q.Exec(
		query,
		u.ID, u.CreatedAt, u.UpdatedAt, u.Email, u.PasswordHash, u.UserStatus, u.UserRole,
	)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
