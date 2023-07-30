package repository

import (
	"github.com/create-go-app/fiber-go-template/database"
)

type BookRepositoryDB struct {
	DB database.DBConn
}

func NewBookRepository(db database.DBConn) BookRepository {
	return &BookRepositoryDB{
		DB: db,
	}
}

type BookRepository interface {
}
