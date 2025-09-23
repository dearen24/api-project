package handlers

import (
	"database/sql"
	"go-fiber-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT id, username, email, password_hash, first_name, last_name, is_active, created_at, updated_at FROM users")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		var users []models.User
		for rows.Next() {
			var user models.User
			if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password_hash, &user.First_name, &user.Last_name, &user.Is_active, &user.Created_at, &user.Updated_at); err != nil {
				return err
			}
			users = append(users, user)
		}

		return c.JSON(users)
	}
}

func GetUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idUser := c.Params("id")
		var user models.User
		err := db.QueryRow("SELECT id, username, email, password_hash, first_name, last_name, is_active, created_at, updated_at FROM users WHERE id = ?", idUser).Scan(&user.Id, &user.Username, &user.Email, &user.Password_hash, &user.First_name, &user.Last_name, &user.Is_active, &user.Created_at, &user.Updated_at)
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		} else if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(user)
	}
}

func CreateUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var newUser struct {
			Username     string `json:"username" binding:"required,alphanum"`
			Password_hash string `json:"password_hash" binding:"required"`
			Email        string `json:"email" binding:"required,email"`
			First_name   string `json:"first_name" binding:"required"`
			Last_name    string `json:"last_name" binding:"required"`
			Is_active		int   `json:"isActive" binding:"required"`
		}

		if err := c.BodyParser(&newUser); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		res, err := db.Exec("INSERT INTO users (username, password_hash, email, first_name, last_name, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())", 
		newUser.Username, newUser.Password_hash, newUser.Email, newUser.First_name, newUser.Last_name, newUser.Is_active)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error creating user"})
		}

		idUser, err := res.LastInsertId()

		var user models.User
		er := db.QueryRow("SELECT id, username, email, password_hash, first_name, last_name, is_active, created_at, updated_at FROM users WHERE id = ?", idUser).Scan(&user.Id, &user.Username, &user.Password_hash, &user.Email, &user.First_name, &user.Last_name, &user.Is_active, &user.Created_at, &user.Updated_at);
		if er != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error fetching data: " + er.Error()})
		}

		return c.Status(201).JSON(user)
	}
}

func EditUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idUser := c.Params("id")
		var updateUser struct {
			Username     string `json:"username" binding:"required,alphanum"`
			Password_hash string `json:"password_hash" binding:"required"`
			Email        string `json:"email" binding:"required,email"`
			First_name   string `json:"first_name" binding:"required"`
			Last_name    string `json:"last_name" binding:"required"`
			Is_active		int   `json:"isActive" binding:"required"`
		}
		if err := c.BodyParser(&updateUser); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		_, err := db.Exec("UPDATE users SET username = ?, password_hash = ?, email = ?, first_name = ?, last_name = ?, is_active = ?, updated_at = NOW() WHERE id = ?",
			updateUser.Username, updateUser.Password_hash, updateUser.Email, updateUser.First_name, updateUser.Last_name, updateUser.Is_active, idUser)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error updating user"})
		}

		var user models.User
		er := db.QueryRow("SELECT id, username, email, password_hash, first_name, last_name, is_active, created_at, updated_at FROM users WHERE id = ?", idUser).Scan(&user.Id, &user.Username, &user.Password_hash, &user.Email, &user.First_name, &user.Last_name, &user.Is_active, &user.Created_at, &user.Updated_at);
		if er != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error fetching data: " + er.Error()})
		}

		return c.JSON(user)
	}
}

func DeleteUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idUser := c.Params("id")
		_, err := db.Exec("DELETE FROM users WHERE id = ?", idUser)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error deleting user"})
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "User deleted successfully",
		})
	}
}
