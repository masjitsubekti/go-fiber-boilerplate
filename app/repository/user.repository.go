package repository

import (
	"fmt"

	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/database"
)

type UserRepositoryDB struct {
	DB database.DBConn
}

func NewUserRepository(db database.DBConn) UserRepository {
	return &UserRepositoryDB{
		DB: db,
	}
}

type UserRepository interface {
	GetUserByID(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	CreateUser(u *models.User) error
}

// GetUserByID query for getting one User by given ID.
func (r *UserRepositoryDB) GetUserByID(id string) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Define query string.
	query := `SELECT * FROM users WHERE id = $1`

	// Send query to database.
	err := r.DB.Query().Get(&user, query, id)
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// GetUserByEmail query for getting one User by given Email.
func (r *UserRepositoryDB) GetUserByEmail(email string) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Define query string.
	query := `SELECT * FROM users WHERE email = $1`

	// Send query to database.
	err := r.DB.Query().Get(&user, query, email)
	if err != nil {
		// Return empty object and error.
		fmt.Println(err)
		return user, err
	}

	// Return query result.
	return user, nil
}

// CreateUser query for creating a new user by given email and password hash.
func (r *UserRepositoryDB) CreateUser(u *models.User) error {
	// Define query string.
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database.
	_, err := r.DB.Query().Exec(
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
