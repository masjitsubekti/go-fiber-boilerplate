package services

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/app/repository"
)

type UserService interface {
	GetUserByID(id string) (user models.User, err error)
	GetUserByEmail(email string) (user models.User, err error)
	CreateUser(u *models.User) error
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: repository,
	}
}

func (r *UserServiceImpl) GetUserByID(id string) (user models.User, err error) {
	return r.UserRepository.GetUserByID(id)
}

func (r *UserServiceImpl) GetUserByEmail(email string) (user models.User, err error) {
	return r.UserRepository.GetUserByEmail(email)
}

func (r *UserServiceImpl) CreateUser(u *models.User) error {
	return r.UserRepository.CreateUser(u)
}
