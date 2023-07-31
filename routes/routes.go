package routes

import (
	"github.com/create-go-app/fiber-go-template/app/middleware"
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

func SetupRoutes(a *fiber.App, injection Injection) {
	SwaggerRoute(a)
	// Create routes group.
	route := a.Group("/api/v1")

	// AUTH
	userController := injection.AuthController
	route.Post("/user/sign/up", userController.UserSignUp)
	route.Post("/user/sign/in", userController.UserSignIn)
	route.Post("/user/sign/out", middleware.JWTProtected(), userController.UserSignOut)
	route.Post("/token/renew", middleware.JWTProtected(), userController.RenewTokens)

	// BOOK
	bookController := injection.BookController
	route.Get("/books", bookController.GetBooks)
	route.Get("/book/:id", bookController.GetBook)
	route.Post("/book", middleware.JWTProtected(), bookController.CreateBook)
	route.Put("/book", middleware.JWTProtected(), bookController.UpdateBook)
	route.Delete("/book", middleware.JWTProtected(), bookController.DeleteBook)
}

func SwaggerRoute(a *fiber.App) {
	route := a.Group("/swagger")
	route.Get("*", swagger.HandlerDefault)
}

func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "sorry, endpoint is not found",
			})
		},
	)
}
