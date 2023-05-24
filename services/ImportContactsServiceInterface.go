package services

import (
	"github.com/gin-gonic/gin"
)

type ImportContactsServiceInterface interface {
	Insert(c *gin.Context)
}
