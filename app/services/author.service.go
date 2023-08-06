package services

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/database"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type AuthorService interface {
	GetAll() (author []models.Author, err error)
	FindByID(id uuid.UUID) (author models.Author, err error)
	Create(req models.AuthorRequest) (res models.Author, err error)
	Update(id uuid.UUID, req models.AuthorRequest) (res models.Author, err error)
	Delete(id uuid.UUID) (err error)
}

type AuthorServiceImpl struct {
	DB database.DBConn
}

func NewAuthorService(db database.DBConn) *AuthorServiceImpl {
	return &AuthorServiceImpl{
		DB: db,
	}
}

func (s *AuthorServiceImpl) GetAll() (res []models.Author, err error) {
	var author models.Author
	err = s.DB.Orm().Table(author.TableName()).
		Select("id", "name", "address").
		Where("deleted_at is null").
		Order("name asc").Scan(&res).Error
	if res == nil {
		return make([]models.Author, 0), nil
	}
	return
}

func (s *AuthorServiceImpl) FindByID(id uuid.UUID) (author models.Author, err error) {
	err = s.DB.Orm().First(&author, "id=?", id).Error
	return
}

func (s *AuthorServiceImpl) Create(req models.AuthorRequest) (res models.Author, err error) {
	res.BindFromRequest(req)
	err = s.DB.Orm().Create(&res).Error
	if err != nil {
		return models.Author{}, err
	}

	return
}

func (s *AuthorServiceImpl) Update(id uuid.UUID, req models.AuthorRequest) (res models.Author, err error) {
	var author models.Author
	err = s.DB.Orm().First(&res, "id=?", id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	author.BindFromRequest(req)
	err = s.DB.Orm().Save(&author).Error
	if err != nil {
		return
	}

	return author, nil
}

func (s *AuthorServiceImpl) Delete(id uuid.UUID) (err error) {
	var author models.Author
	err = s.DB.Orm().First(&author, "id=?", id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	author.SoftDelete()
	err = s.DB.Orm().Save(&author).Error
	if err != nil {
		return
	}
	return nil
}
