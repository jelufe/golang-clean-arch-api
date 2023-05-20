package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	Username     *string            `json:"username" validate:"required"`
	Password     *string            `json:"password" validate:"required"`
	UserType     *string            `json:"user_type" validate:"required"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
}
