package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jelufe/golang-clean-arch-api/models"
)

// importContacts	godoc
// @Sumary import Contacts
// @Description Save contacts data in database
// @Param users body models.ImportContactsRequest true "Contacts"
// @Produce application/json
// @Tags contacts
// @Success 200
// @Failure 400
// @Failure 500
// @Router /contacts [post]
func ImportContacts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var importContactsRequest models.ImportContactsRequest
		err := c.BindJSON(&importContactsRequest)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}
