package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jelufe/golang-clean-arch-api/database"
	"github.com/jelufe/golang-clean-arch-api/models"
)

func NewContactMacapaRepository() ContactMacapaRepositoryInterface {
	return &ContactMacapaRepository{
		MysqlDb: database.MysqlDb,
	}
}

type ContactMacapaRepository struct {
	MysqlDb *sql.DB
}

func (repositories *ContactMacapaRepository) Insert(contacts []models.Contact) (rowsAffected int64, err error) {
	sqlStr := "INSERT INTO contacts (nome, celular) VALUES "

	for _, row := range contacts {
		sqlStr += fmt.Sprintf("('%v', '%v'),", row.Name, row.Cellphone)
	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]

	res, execError := repositories.MysqlDb.Exec(sqlStr)

	if execError != nil {
		return 0, execError
	}

	rowsAffected, _ = res.RowsAffected()

	return rowsAffected, execError
}
