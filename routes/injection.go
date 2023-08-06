package routes

import (
	"github.com/create-go-app/fiber-go-template/app/controllers"
	"github.com/create-go-app/fiber-go-template/app/repository"
	"github.com/create-go-app/fiber-go-template/app/services"
	"github.com/create-go-app/fiber-go-template/database"
)

type Injection struct {
	AuthController   controllers.AuthController
	BookController   controllers.BookController
	AuthorController controllers.AuthorController
}

// Define Dependency Injection
func CallDependenciesInjection() Injection {
	DbConnect, _ := database.NewDBConnection()
	// Auth
	userRepository := repository.NewUserRepository(DbConnect)
	userService := services.NewUserService(userRepository)
	authController := controllers.NewAuthController(userService)
	// Book
	bookRepository := repository.NewBookRepository(DbConnect)
	bookService := services.NewBookService(bookRepository)
	bookController := controllers.NewBookController(bookService)
	// Author
	authorService := services.NewAuthorService(DbConnect)
	authorController := controllers.NewAuthorController(authorService)

	return Injection{
		AuthController:   authController,
		BookController:   bookController,
		AuthorController: authorController,
	}
}
