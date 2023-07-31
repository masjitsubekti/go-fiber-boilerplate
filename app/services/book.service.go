package services

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/app/repository"
	"github.com/gofrs/uuid"
)

type BookService interface {
	GetBooks() ([]models.Book, error)
	GetBooksByAuthor(author string) ([]models.Book, error)
	GetBook(id uuid.UUID) (models.Book, error)
	CreateBook(b *models.Book) error
	UpdateBook(id uuid.UUID, b *models.Book) error
	DeleteBook(id uuid.UUID) error
}

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookService(repository repository.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{
		BookRepository: repository,
	}
}

func (s *BookServiceImpl) GetBooks() ([]models.Book, error) {
	return s.BookRepository.GetBooks()
}

func (s *BookServiceImpl) GetBooksByAuthor(author string) ([]models.Book, error) {
	return s.BookRepository.GetBooksByAuthor(author)
}

func (s *BookServiceImpl) GetBook(id uuid.UUID) (models.Book, error) {
	return s.BookRepository.GetBook(id)
}

func (s *BookServiceImpl) CreateBook(b *models.Book) error {
	return s.BookRepository.CreateBook(b)
}

func (s *BookServiceImpl) UpdateBook(id uuid.UUID, b *models.Book) error {
	return s.BookRepository.UpdateBook(id, b)
}

func (s *BookServiceImpl) DeleteBook(id uuid.UUID) error {
	return s.BookRepository.DeleteBook(id)
}
