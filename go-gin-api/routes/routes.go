package routes

import (
	"github.com/gin-gonic/gin"
	"example/go-gin-api/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	// r.GET("/users", handlers.GetUsers)
	// r.GET("/users/:id", handlers.GetUser)
	// r.POST("/users", handlers.CreateUser)

	protected := r.Group("/api")
	protected.Use(handlers.AuthRequired())
	{
		protected.GET("/users", handlers.GetUsers)
		protected.GET("/users/:id", handlers.GetUser)
		protected.POST("/users", handlers.CreateUser)
	}
}
