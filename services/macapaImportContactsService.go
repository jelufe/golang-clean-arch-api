package services

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jelufe/golang-clean-arch-api/database"
	"github.com/jelufe/golang-clean-arch-api/models"
)

var MysqlDb *sql.DB = database.MysqlDb

func MacapaImportContacts(importContactsRequest models.ImportContactsRequest) (rowsAffected int64, err error) {
	sqlStr := "INSERT INTO contacts (nome, celular) VALUES "

	var re = regexp.MustCompile(`[^0-9.]`)

	for _, row := range *importContactsRequest.Contacts {
		name := strings.ToUpper(row.Name)
		cellphone := re.ReplaceAllString(row.Cellphone, "")
		cellphoneMasked := "+" + cellphone[0:2] + " (" + cellphone[2:4] + ") " + cellphone[4:9] + "-" + cellphone[9:13]
		sqlStr += fmt.Sprintf("('%v', '%v'),", name, cellphoneMasked)
	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]

	res, execError := MysqlDb.Exec(sqlStr)

	if execError != nil {
		return 0, execError
	}

	rowsAffected, _ = res.RowsAffected()

	return rowsAffected, execError
}
