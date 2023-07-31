package bootstrap

import (
	"os"

	"github.com/create-go-app/fiber-go-template/app/middleware"
	"github.com/create-go-app/fiber-go-template/config"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/create-go-app/fiber-go-template/routes"
	"github.com/gofiber/fiber/v2"
)

func Boot() {
	// Define Fiber config.
	config := config.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app)

	// Dependencies Injection
	injection := routes.CallDependenciesInjection()

	// Routes.
	routes.SetupRoutes(app, injection)
	routes.NotFoundRoute(app)

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
