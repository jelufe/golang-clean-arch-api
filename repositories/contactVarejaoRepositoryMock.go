package repositories

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jelufe/golang-clean-arch-api/models"
)

func NewContactVarejaoRepositoryMock() ContactVarejaoRepositoryInterface {
	return &ContactVarejaoRepositoryMock{}
}

type ContactVarejaoRepositoryMock struct{}

func (repositories *ContactVarejaoRepositoryMock) Insert(contacts []models.Contact) (rowsAffected int64, err error) {
	return 1, nil
}
