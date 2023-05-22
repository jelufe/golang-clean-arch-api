package services

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jelufe/golang-clean-arch-api/database"
	"github.com/jelufe/golang-clean-arch-api/models"
)

var MysqlDb *sql.DB = database.MysqlDb

func MacapaImportContacts(importContactsRequest models.ImportContactsRequest) (rowsAffected int64, err error) {
	sqlStr := "INSERT INTO contacts (nome, celular) VALUES "

	for _, row := range *importContactsRequest.Contacts {
		name := row.Name
		cellphone := row.Cellphone
		sqlStr += fmt.Sprintf("('%v', '%v'),", name, cellphone)
	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]

	res, execError := MysqlDb.Exec(sqlStr)

	database.CloseMysqlDb()

	if execError != nil {
		return 0, execError
	}

	rowsAffected, _ = res.RowsAffected()

	return rowsAffected, execError
}
