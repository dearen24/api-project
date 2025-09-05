package handlers

import (
	"net/http"
	"database/sql"
	"time"
	"os"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"
	"example/go-gin-api/models"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
    DB *sql.DB
}

func (h *UserHandler) Login(c *gin.Context) {
	var input struct {
		Username    string `json:"username" binding:"required,alphanum"`
		Password_hash string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := h.DB.QueryRow("SELECT id, username, password_hash FROM users WHERE username = ?", input.Username).Scan(&user.Id, &user.Username, &user.Password_hash); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentialssssss"})
		return
	}

	// Compare password
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	// 	return
	// }

	if user.Password_hash != input.Password_hash {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Create JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	rows, err := h.DB.Query("SELECT * FROM users")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
        return
    }
    defer rows.Close()

    var users []map[string]interface{}
    for rows.Next() {
        var id int
        var username string
		var password_hash string
		var email string
		var first_name string
		var last_name string
		var is_active int
		var created_at string
		var updated_at string
        rows.Scan(&id, &username, &password_hash, &email, &first_name, &last_name, &is_active, &created_at, &updated_at)
        users = append(users, map[string]interface{}{"id": id, "username": username, "password_hash": password_hash, "email": email, 
		"first_name": first_name, "last_name": last_name, "is_active": is_active, "created_at": created_at, "updated_at": updated_at})
    }
    c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idUser, _ := strconv.Atoi(c.Param("id"))

	var (
        id        	  int
        username      string
        password_hash string
        email         string
        first_name    string
        last_name     string
        is_active     int
        created_at    string
        updated_at    string
    )

	err := h.DB.QueryRow("SELECT * FROM users WHERE id = ?", idUser).Scan(&id, &username, &password_hash, &email, &first_name, &last_name, &is_active, &created_at, &updated_at);
	if err == sql.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
        return
    }

	user := map[string]interface{}{"id": id, "username": username, "password_hash": password_hash, "email": email,
	"first_name": first_name, "last_name": last_name, "is_active": is_active, "created_at": created_at, "updated_at": updated_at}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser struct {
        Username     string `json:"username" binding:"required,alphanum"`
        Password_hash string `json:"password" binding:"required"`
		Email        string `json:"email" binding:"required,email"`
		First_name   string `json:"firstName" binding:"required"`
		Last_name    string `json:"lastName" binding:"required"`
		Is_active		int   `json:"isActive" binding:"required"`
    }

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	res, err := h.DB.Exec(`INSERT INTO users (username, password_hash, email, first_name, last_name, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())`,
		newUser.Username, newUser.Password_hash, newUser.Email, newUser.First_name, newUser.Last_name, newUser.Is_active)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		fmt.Println(err)
        return
    }

	idUser, err := res.LastInsertId()

	var (
        id        	  int
        username      string
        password_hash string
        email         string
        first_name    string
        last_name     string
        is_active     int
        created_at    string
        updated_at    string
    )

	er := h.DB.QueryRow("SELECT * FROM users WHERE id = ?", idUser).Scan(&id, &username, &password_hash, &email, &first_name, &last_name, &is_active, &created_at, &updated_at);
	if er == sql.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    } else if er != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
        return
    }

	user := map[string]interface{}{"id": id, "username": username, "password_hash": password_hash, "email": email,
	"first_name": first_name, "last_name": last_name, "is_active": is_active, "created_at": created_at, "updated_at": updated_at}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) EditUser(c *gin.Context) {
	idUser, _ := strconv.Atoi(c.Param("id"))

	var data struct {
        Username     string `json:"username" binding:"required,alphanum"`
        Password_hash string `json:"password" binding:"required"`
		Email        string `json:"email" binding:"required,email"`
		First_name   string `json:"firstName" binding:"required"`
		Last_name    string `json:"lastName" binding:"required"`
		Is_active		*int   `json:"isActive" binding:"required"`
    }

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorrrr": err.Error()})
		return
	}
	
	_, err := h.DB.Exec(`UPDATE users SET username = ?, email = ?, password_hash = ?, first_name = ?, last_name = ?, is_active = ?, updated_at = NOW() WHERE id=?`,
		data.Username, data.Email, data.Password_hash, data.First_name, data.Last_name, data.Is_active, idUser)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		fmt.Println(err)
        return
    }

	var (
        id        	  int
        username      string
        password_hash string
        email         string
        first_name    string
        last_name     string
        is_active     int
        created_at    string
        updated_at    string
    )

	er := h.DB.QueryRow("SELECT * FROM users WHERE id = ?", idUser).Scan(&id, &username, &password_hash, &email, &first_name, &last_name, &is_active, &created_at, &updated_at);
	if er == sql.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    } else if er != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
        return
    }

	user := map[string]interface{}{"id": id, "username": username, "password_hash": password_hash, "email": email,
	"first_name": first_name, "last_name": last_name, "is_active": is_active, "created_at": created_at, "updated_at": updated_at}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idUser, _ := strconv.Atoi(c.Param("id"))

	_, err := h.DB.Exec("DELETE FROM users WHERE id = ?", idUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
