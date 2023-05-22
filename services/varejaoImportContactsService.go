package services

import (
	"database/sql"
	"fmt"

	"github.com/jelufe/golang-clean-arch-api/database"
	"github.com/jelufe/golang-clean-arch-api/models"
	_ "github.com/lib/pq"
)

var PostgresDb *sql.DB = database.PostgresDb

func VarejaoImportContacts(importContactsRequest models.ImportContactsRequest) (rowsAffected int64, err error) {
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
		return 0, execError
	}

	rowsAffected, _ = res.RowsAffected()

	return rowsAffected, execError
}
