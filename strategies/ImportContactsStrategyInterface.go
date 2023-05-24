package strategies

import "github.com/jelufe/golang-clean-arch-api/models"

type ImportContactsStrategyInterface interface {
	Insert(importContactsRequest models.ImportContactsRequest) (rowsAffected int64, err error)
}
