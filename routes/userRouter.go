package routes

import (
	controller "github.com/jelufe/golang-clean-arch-api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jelufe/golang-clean-arch-api/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:id", controller.GetUser())
}
