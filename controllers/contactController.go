package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jelufe/golang-clean-arch-api/services"
)

// importContacts	godoc
// @Sumary import Contacts
// @Description Save contacts data in database
// @Param users body models.ImportContactsRequest true "Contacts"
// @Produce application/json
// @Tags contacts
// @Success 200 {object} models.ImportContactsResponse
// @Failure 400
// @Failure 500
// @Router /contacts [post]
// @Security BearerAuth
func ImportContacts() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.NewImportContactsService().Insert(c)
	}
}
