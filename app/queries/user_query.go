package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/koddr/tutorial-go-fiber-rest-api/app/models"
)

// UserQueries struct for queries_v1 from User model.
type UserQueries struct {
	*sqlx.DB
}

// GetUsers method for getting all Users.
func (q *UserQueries) GetUsers() ([]models.User, error) {
	// Define Users variable.
	var Users []models.User

	println("Created list of users")
	println(Users)

	// Define query string.
	query := `SELECT * FROM users`
	println(query)

	// Send query to database.
	err := q.Get(&Users, query)
	println(err.Error())
	println(err)
	println(Users)
	if err != nil {
		// Return empty object and error.
		return Users, err
	}

	// Return query result.
	return Users, nil
}

// GetUser method for getting one User by given ID.
func (q *UserQueries) GetUser(id uuid.UUID) (models.User, error) {
	// Define User variable.
	User := models.User{}

	// Define query string.
	query := `SELECT * FROM users WHERE id = $1`

	// Send query to database.
	err := q.Get(&User, query, id)
	if err != nil {
		// Return empty object and error.
		return User, err
	}

	// Return query result.
	return User, nil
}

// CreateUser method for creating User by given User object.
func (q *UserQueries) CreateUser(b *models.User) error {
	// Define query string.
	query := `INSERT INTO users VALUES ($1, $2, $3, $4)`

	// Send query to database.
	_, err := q.Exec(query, b.Id, b.Login, b.Password, b.AcsLevel)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// UpdateUser method for updating User by given User object.
func (q *UserQueries) UpdateUser(id uuid.UUID, b *models.User) error {
	// Define query string.
	query := `UPDATE users SET login = $2, password = $3, acs_level = $4 WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id, b.Login, b.Password, b.AcsLevel)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// DeleteUser method for delete User by given ID.
func (q *UserQueries) DeleteUser(id uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM Users WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
