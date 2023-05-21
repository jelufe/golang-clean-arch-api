package models

type ImportContactsRequest struct {
	Contacts *[]Contact `json:"contacts"`
}
