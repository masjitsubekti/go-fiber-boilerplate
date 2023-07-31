package repository

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/database"
	"github.com/gofrs/uuid"
)

type BookRepository interface {
	GetBooks() ([]models.Book, error)
	GetBooksByAuthor(author string) ([]models.Book, error)
	GetBook(id uuid.UUID) (models.Book, error)
	CreateBook(b *models.Book) error
	UpdateBook(id uuid.UUID, b *models.Book) error
	DeleteBook(id uuid.UUID) error
}

type BookRepositoryDB struct {
	DB database.DBConn
}

func NewBookRepository(db database.DBConn) BookRepository {
	return &BookRepositoryDB{
		DB: db,
	}
}

// GetBooks method for getting all books.
func (r *BookRepositoryDB) GetBooks() ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Define query string.
	query := `SELECT * FROM books`

	// Send query to database.
	err := r.DB.Query().Select(&books, query)
	if err != nil {
		// Return empty object and error.
		return books, err
	}

	// Return query result.
	return books, nil
}

// GetBooksByAuthor method for getting all books by given author.
func (r *BookRepositoryDB) GetBooksByAuthor(author string) ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Define query string.
	query := `SELECT * FROM books WHERE author = $1`

	// Send query to database.
	err := r.DB.Query().Get(&books, query, author)
	if err != nil {
		// Return empty object and error.
		return books, err
	}

	// Return query result.
	return books, nil
}

// GetBook method for getting one book by given ID.
func (r *BookRepositoryDB) GetBook(id uuid.UUID) (models.Book, error) {
	// Define book variable.
	book := models.Book{}

	// Define query string.
	query := `SELECT * FROM books WHERE id = $1`

	// Send query to database.
	err := r.DB.Query().Get(&book, query, id)
	if err != nil {
		// Return empty object and error.
		return book, err
	}

	// Return query result.
	return book, nil
}

// CreateBook method for creating book by given Book object.
func (r *BookRepositoryDB) CreateBook(b *models.Book) error {
	// Define query string.
	query := `INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Send query to database.
	_, err := r.DB.Query().Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.UserID, b.Title, b.Author, b.BookStatus, b.BookAttrs)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// UpdateBook method for updating book by given Book object.
func (r *BookRepositoryDB) UpdateBook(id uuid.UUID, b *models.Book) error {
	// Define query string.
	query := `UPDATE books SET updated_at = $2, title = $3, author = $4, book_status = $5, book_attrs = $6 WHERE id = $1`

	// Send query to database.
	_, err := r.DB.Query().Exec(query, id, b.UpdatedAt, b.Title, b.Author, b.BookStatus, b.BookAttrs)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// DeleteBook method for delete book by given ID.
func (r *BookRepositoryDB) DeleteBook(id uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM books WHERE id = $1`

	// Send query to database.
	_, err := r.DB.Query().Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
