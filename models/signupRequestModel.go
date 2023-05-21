package models

type SignupRequest struct {
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password" validate:"required"`
	UserType *string `json:"user_type" validate:"required"`
}
