package strategies

import (
	"testing"

	"github.com/jelufe/golang-clean-arch-api/models"
	"github.com/jelufe/golang-clean-arch-api/repositories"
)

func TestMacapaInsert(t *testing.T) {
	importContactsMacapaStrategy := ImportContactsMacapaStrategy{MacapaRepository: repositories.NewContactMacapaRepositoryMock()}
	contacts := []models.Contact{}
	contacts = append(contacts, models.Contact{Name: "Teste", Cellphone: "9999999999999"})
	importContactsRequest := models.ImportContactsRequest{Contacts: contacts}

	got, _ := importContactsMacapaStrategy.Insert(importContactsRequest)
	var want int64 = 1

	if got != want {
		t.Errorf("got rowsAffected %v, wanted %v", got, want)
	}
}
