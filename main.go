package main

import (
	"os"

	_ "github.com/jelufe/golang-clean-arch-api/docs"

	routes "github.com/jelufe/golang-clean-arch-api/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Go Service API
// @version 1.0
// @description A Service API in Go using Gin framework
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:9000
// @BasePath /
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.ContactRoutes(router)

	router.Run(":" + port)
}
