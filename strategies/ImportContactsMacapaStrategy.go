package strategies

import (
	"regexp"
	"strings"

	"github.com/jelufe/golang-clean-arch-api/models"
	"github.com/jelufe/golang-clean-arch-api/repositories"
)

func NewImportContactsMacapaStrategy() ImportContactsStrategyInterface {
	return &ImportContactsMacapaStrategy{
		MacapaRepository: repositories.NewContactMacapaRepository(),
	}
}

type ImportContactsMacapaStrategy struct {
	MacapaRepository repositories.ContactMacapaRepositoryInterface
}

func (strategies ImportContactsMacapaStrategy) Insert(importContactsRequest models.ImportContactsRequest) (rowsAffected int64, err error) {
	var re = regexp.MustCompile(`[^0-9.]`)

	contacts := importContactsRequest.Contacts

	for i, row := range importContactsRequest.Contacts {
		row.Name = strings.ToUpper(row.Name)
		cellphone := re.ReplaceAllString(row.Cellphone, "")
		row.Cellphone = "+" + cellphone[0:2] + " (" + cellphone[2:4] + ") " + cellphone[4:9] + "-" + cellphone[9:13]
		contacts[i] = row
	}

	rowsAffected, execError := strategies.MacapaRepository.Insert(contacts)

	return rowsAffected, execError
}
