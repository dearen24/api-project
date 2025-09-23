package routes

import (
	"database/sql"
	"go-fiber-api/handlers"
	"go-fiber-api/auth"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	api := app.Group("/api")

	// Public routes
	api.Post("/login", handlers.Login(db))

	// JWT Middleware for protected routes
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: auth.SecretKey(),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or malformed JWT",
			})
		},
	}))

	// Protected routes
	api.Get("/users", handlers.GetUsers(db))
	api.Get("/users/:id", handlers.GetUser(db))
	api.Post("/users", handlers.CreateUser(db))
	api.Put("/users/:id", handlers.EditUser(db))
	api.Delete("/users/:id", handlers.DeleteUser(db))
}
