package strategies

import (
	"regexp"

	"github.com/jelufe/golang-clean-arch-api/models"
	"github.com/jelufe/golang-clean-arch-api/repositories"
)

func NewImportContactsVarejaoStrategy() ImportContactsStrategyInterface {
	return &ImportContactsVarejaoStrategy{
		VarejaoRepository: repositories.NewContactVarejaoRepository(),
	}
}

type ImportContactsVarejaoStrategy struct {
	VarejaoRepository repositories.ContactVarejaoRepositoryInterface
}

func (strategies ImportContactsVarejaoStrategy) Insert(importContactsRequest models.ImportContactsRequest) (rowsAffected int64, err error) {
	var re = regexp.MustCompile(`[^0-9.]`)

	contacts := importContactsRequest.Contacts

	for i, row := range importContactsRequest.Contacts {
		row.Cellphone = re.ReplaceAllString(row.Cellphone, "")
		contacts[i] = row
	}

	rowsAffected, execError := strategies.VarejaoRepository.Insert(importContactsRequest.Contacts)

	return rowsAffected, execError
}
