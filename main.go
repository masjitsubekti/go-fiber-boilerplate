package main

//go:generate go run github.com/swaggo/swag/cmd/swag init

import (
	"github.com/create-go-app/fiber-go-template/bootstrap"

	_ "github.com/create-go-app/fiber-go-template/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title Go Fiber Boilerplate
// @version 1.0
// @description This is an auto-generated API Docs.

// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	bootstrap.Boot()
}
