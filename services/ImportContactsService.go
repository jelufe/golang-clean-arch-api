package services

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/jelufe/golang-clean-arch-api/models"
	"github.com/jelufe/golang-clean-arch-api/strategies"
)

func NewImportContactsService() ImportContactsServiceInterface {
	return &ImportContactsService{
		ImportContactsMacapaStrategy:  strategies.NewImportContactsMacapaStrategy(),
		ImportContactsVarejaoStrategy: strategies.NewImportContactsVarejaoStrategy(),
	}
}

type ImportContactsService struct {
	ImportContactsMacapaStrategy  strategies.ImportContactsStrategyInterface
	ImportContactsVarejaoStrategy strategies.ImportContactsStrategyInterface
}

func (services ImportContactsService) Insert(c *gin.Context) {
	var importContactsRequest models.ImportContactsRequest
	err := c.BindJSON(&importContactsRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var re = regexp.MustCompile(`[^0-9.]`)
	for _, row := range importContactsRequest.Contacts {
		cellphone := re.ReplaceAllString(row.Cellphone, "")
		if len(cellphone) != 13 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "All cellphones must be 13 numbers"})
			return
		}
	}

	userType := c.GetString("user_type")
	var rowsAffected int64
	var execError error

	if userType == "VAREJAO" {
		rowsAffected, execError = services.ImportContactsVarejaoStrategy.Insert(importContactsRequest)
	} else if userType == "MACAPA" {
		rowsAffected, execError = services.ImportContactsMacapaStrategy.Insert(importContactsRequest)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unmapped import for this user"})
	}

	if execError != nil {
		c.JSON(http.StatusInternalServerError, execError.Error())
		return
	}

	c.JSON(http.StatusOK, models.ImportContactsResponse{RowsAffected: rowsAffected})
}
