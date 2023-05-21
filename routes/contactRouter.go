package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jelufe/golang-clean-arch-api/controllers"
	"github.com/jelufe/golang-clean-arch-api/middleware"
)

func ContactRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/contacts", controller.ImportContacts())
}
