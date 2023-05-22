package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jelufe/golang-clean-arch-api/database"
	"github.com/jelufe/golang-clean-arch-api/models"
	_ "github.com/lib/pq"
)

var PostgresDb *sql.DB = database.PostgresDb

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

		sqlStr := "INSERT INTO contacts (nome, celular) VALUES "

		for _, row := range *importContactsRequest.Contacts {
			name := row.Name
			cellphone := row.Cellphone
			sqlStr += fmt.Sprintf("('%v', '%v'),", name, cellphone)
		}

		sqlStr = sqlStr[0 : len(sqlStr)-1]

		res, execError := PostgresDb.Exec(sqlStr)

		database.ClosePostgresDb()

		if execError != nil {
			c.JSON(http.StatusInternalServerError, execError.Error())
			return
		}

		rowsAffected, _ := res.RowsAffected()

		c.JSON(http.StatusOK, rowsAffected)
	}
}
