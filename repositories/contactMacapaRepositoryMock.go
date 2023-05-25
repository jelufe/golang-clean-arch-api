package repositories

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jelufe/golang-clean-arch-api/models"
)

func NewContactMacapaRepositoryMock() ContactMacapaRepositoryInterface {
	return &ContactMacapaRepositoryMock{}
}

type ContactMacapaRepositoryMock struct{}

func (repositories *ContactMacapaRepositoryMock) Insert(contacts []models.Contact) (rowsAffected int64, err error) {
	return 1, nil
}
