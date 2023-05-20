package main

import (
	"os"

	routes "github.com/jelufe/golang-clean-arch-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api"})
	})

	router.Run(":" + port)
}
