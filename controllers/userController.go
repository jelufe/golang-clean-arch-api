package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/jelufe/golang-clean-arch-api/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup()

func Login()

func GetUsers()

func GetUser()