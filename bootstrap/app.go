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
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
