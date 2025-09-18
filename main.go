// @title Training KBTG Backend API
// @version 1.0
// @description This is a training backend API with authentication
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"log"
	"temp-backend-at-kbtg/database"
	_ "temp-backend-at-kbtg/docs"
	"temp-backend-at-kbtg/handlers"
	"temp-backend-at-kbtg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// HelloWorld godoc
// @Summary Get hello world message
// @Description Get a simple hello world message
// @Tags General
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func helloWorld(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "hello world",
	})
}

// ProtectedRoute godoc
// @Summary Protected route example
// @Description Example of a protected route that requires authentication
// @Tags General
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /protected [get]
func protectedRoute(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	email := c.Locals("email")

	return c.JSON(fiber.Map{
		"message": "This is a protected route",
		"user_id": userID,
		"email":   email,
	})
}

func main() {
	// Connect to database
	database.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		AppName: "Training KBTG Backend API v1.0.0",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	}))

	// Swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Routes
	app.Get("/", helloWorld)

	// Auth routes
	auth := app.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Protected routes
	app.Get("/protected", middleware.JWTMiddleware(), protectedRoute)

	// Profile routes
	profile := app.Group("/profile", middleware.JWTMiddleware())
	profile.Get("/", handlers.GetProfile)
	profile.Put("/", handlers.UpdateProfile)
	profile.Get("/membership", handlers.GetMembershipInfo)

	// Start server on port 3000
	log.Printf("Server starting on port 3000...")
	log.Printf("Swagger documentation available at http://localhost:3000/swagger/")
	log.Fatal(app.Listen(":3000"))
}