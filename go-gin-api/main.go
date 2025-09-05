package main

import (
	"github.com/gin-gonic/gin"
	"example/go-gin-api/routes"
	"example/go-gin-api/db"
	"example/go-gin-api/handlers"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load("../.env")
    if err != nil {
        log.Fatal("Error loading .env file:", err)
    }

	db.InitDB();
	defer db.DB.Close();
	
	userHandler := &handlers.UserHandler{DB: db.DB}

	router := gin.Default()
	routes.RegisterUserRoutes(router, userHandler)
	router.Run(":3000")
}
