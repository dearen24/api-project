package routes

import (
	"github.com/gin-gonic/gin"
	"example/go-gin-api/handlers"
)

func RegisterUserRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	// r.GET("/users", handlers.GetUsers)
	// r.GET("/users/:id", handlers.GetUser)
	// r.POST("/users", handlers.CreateUser)
	r.POST("api/login", userHandler.Login)

	protected := r.Group("/api")
	protected.Use(handlers.AuthRequired())
	{
		protected.GET("/users", userHandler.GetUsers)
		protected.GET("/users/:id", userHandler.GetUser)
		protected.POST("/users", userHandler.CreateUser)
		protected.PUT("/users/:id", userHandler.EditUser)
		protected.DELETE("/users/:id", userHandler.DeleteUser)
	}
}
