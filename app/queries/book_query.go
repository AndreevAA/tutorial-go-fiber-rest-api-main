package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/koddr/tutorial-go-fiber-rest-api/app/models"
)

// BookQueries struct for queries from Book model.
type BookQueries struct {
	*sqlx.DB
}

// GetBooks method for getting all books.
func (q *BookQueries) GetBooks() ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Define query string.
	query := `SELECT * FROM books`

	println("getted query")

	// Send query to database.
	err := q.Select(&books, query)

	println(books)
	if err != nil {
		// Return empty object and error.
		return books, err
	}

	// Return query result.
	return books, nil
}

// GetBook method for getting one book by given ID.
func (q *BookQueries) GetBook(id uuid.UUID) (models.Book, error) {
	// Define book variable.
	book := models.Book{}

	// Define query string.
	query := `SELECT * FROM books WHERE id = $1`

	// Send query to database.
	err := q.Get(&book, query, id)
	if err != nil {
		// Return empty object and error.
		return book, err
	}

	// Return query result.
	return book, nil
}

// CreateBook method for creating book by given Book object.
func (q *BookQueries) CreateBook(b *models.Book) error {
	// Define query string.
	query := `INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database.
	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.Title, b.Author, b.BookStatus, b.BookAttrs)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// UpdateBook method for updating book by given Book object.
func (q *BookQueries) UpdateBook(id uuid.UUID, b *models.Book) error {
	// Define query string.
	query := `UPDATE books SET updated_at = $2, title = $3, author = $4, book_status = $5, book_attrs = $6 WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id, b.UpdatedAt, b.Title, b.Author, b.BookStatus, b.BookAttrs)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// DeleteBook method for delete book by given ID.
func (q *BookQueries) DeleteBook(id uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM books WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}