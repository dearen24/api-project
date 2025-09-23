package main

import (
	"log"
	"go-fiber-api/config"
	"go-fiber-api/db"
	"go-fiber-api/routes"
	"go-fiber-api/auth"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Init auth with secret
	auth.InitAuth(cfg)

	// Connect to DB
	database, err := db.ConnectMySQL(cfg)
	if err != nil {
		log.Fatal("❌ Database connection failed: ", err)
	}
	log.Println("✅ Connected to MySQL successfully!")
	defer database.Close()

	// Setup Fiber
	app := fiber.New()

	// Register routes
	routes.SetupRoutes(app, database)

	// Start server
	log.Fatal(app.Listen(":" + cfg.Port))
}
