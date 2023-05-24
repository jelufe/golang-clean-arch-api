package repositories

import (
	"database/sql"
	"fmt"

	"github.com/jelufe/golang-clean-arch-api/database"
	"github.com/jelufe/golang-clean-arch-api/models"
	_ "github.com/lib/pq"
)

func NewContactVarejaoRepository() ContactVarejaoRepositoryInterface {
	return &ContactVarejaoRepository{
		PostgresDb: database.PostgresDb,
	}
}

type ContactVarejaoRepository struct {
	PostgresDb *sql.DB
}

func (repositories *ContactVarejaoRepository) Insert(contacts []models.Contact) (rowsAffected int64, err error) {
	sqlStr := "INSERT INTO contacts (nome, celular) VALUES "

	for _, row := range contacts {
		sqlStr += fmt.Sprintf("('%v', '%v'),", row.Name, row.Cellphone)
	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]

	res, execError := repositories.PostgresDb.Exec(sqlStr)

	if execError != nil {
		return 0, execError
	}

	rowsAffected, _ = res.RowsAffected()

	return rowsAffected, execError
}
