package models

type Contact struct {
	Name      string `json:"name" validate:"required"`
	Cellphone string `json:"cellphone" validate:"required"`
}
