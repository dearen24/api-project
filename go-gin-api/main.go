package main

import (
	"github.com/gin-gonic/gin"
	"example/go-gin-api/routes"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}
