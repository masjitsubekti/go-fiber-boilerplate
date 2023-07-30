package repository

import "github.com/create-go-app/fiber-go-template/database"

type UserRepositoryDB struct {
	DB database.DBConn
}

func NewUserRepository(db database.DBConn) UserRepository {
	return &BookRepositoryDB{
		DB: db,
	}
}

type UserRepository interface {
}
