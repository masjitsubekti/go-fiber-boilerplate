package routes

import (
	"github.com/create-go-app/fiber-go-template/app/controllers"
	"github.com/create-go-app/fiber-go-template/app/middleware"
	"github.com/create-go-app/fiber-go-template/app/repository"
	"github.com/create-go-app/fiber-go-template/app/services"
	"github.com/create-go-app/fiber-go-template/database"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Define Dependency Injection
	DbConnect, _ := database.NewDBConnection()
	userRepository := repository.NewUserRepository(DbConnect)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewAuthController(userService)

	// Create routes group.
	route := a.Group("/api/v1")

	// AUTH
	route.Post("/user/sign/up", userController.UserSignUp)                              // register a new user
	route.Post("/user/sign/in", userController.UserSignIn)                              // auth, return Access & Refresh tokens
	route.Post("/user/sign/out", middleware.JWTProtected(), userController.UserSignOut) // de-authorization user
	route.Post("/token/renew", middleware.JWTProtected(), userController.RenewTokens)   // renew Access & Refresh tokens

	// BOOK
	route.Get("/books", controllers.GetBooks)                                // get list of all books
	route.Get("/book/:id", controllers.GetBook)                              // get one book by ID
	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook)   // create a new book
	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook)    // update one book by ID
	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook) // delete one book by ID
}
