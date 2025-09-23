package handlers

import (
	"database/sql"
	"go-fiber-api/auth"
	"go-fiber-api/models"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		var user models.User
		err := db.QueryRow("SELECT id, username, password_hash FROM users WHERE username = ?", req.Username).Scan(&user.Id, &user.Username, &user.Password_hash)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
		}

		// Generate token
		token, err := auth.GenerateToken(user.Id)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
		}

		return c.JSON(fiber.Map{"token": token})
	}
}
