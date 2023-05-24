package repositories

import "github.com/jelufe/golang-clean-arch-api/models"

type ContactVarejaoRepositoryInterface interface {
	Insert(contacts []models.Contact) (rowsAffected int64, err error)
}
