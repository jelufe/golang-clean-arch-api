package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jelufe/golang-clean-arch-api/database"
	helper "github.com/jelufe/golang-clean-arch-api/helpers"
	"github.com/jelufe/golang-clean-arch-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("username or password is incorrect")
		check = false
	}

	return check, msg
}

// Signup	godoc
// @Sumary Create user
// @Description Save user data in database
// @Param users body models.SignupRequest true "Signup"
// @Produce application/json
// @Tags users
// @Success 200
// @Failure 400
// @Failure 500
// @Router /users/signup [post]
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var signupRequest models.SignupRequest
		err := c.BindJSON(&signupRequest)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := models.User{Username: signupRequest.Username, Password: signupRequest.Password, UserType: signupRequest.Password}
		validationError := validate.Struct(user)

		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
			return
		}

		count, countError := userCollection.CountDocuments(ctx, bson.M{"username": user.Username})

		defer cancel()

		if countError != nil {
			log.Panic(countError)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error checking if username exists"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "username already exists try using another one"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password
		user.Id = primitive.NewObjectID()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Username, *user.UserType, user.Id)
		user.Token = &token
		user.RefreshToken = &refreshToken

		resultInsertionNumber, insertError := userCollection.InsertOne(ctx, user)

		if insertError != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

// Login godoc
// @Sumary Login user
// @Description Generate access token
// @Param users body models.LoginRequest true "Login"
// @Produce application/json
// @Tags users
// @Success 200
// @Failure 400
// @Failure 500
// @Router /users/login [post]
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var loginRequest models.LoginRequest
		err := c.BindJSON(&loginRequest)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := models.User{Username: loginRequest.Username, Password: loginRequest.Password}
		var foundUser models.User

		findError := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
		defer cancel()

		if findError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "username or password is incorrect"})
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()

		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if foundUser.Username == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		}

		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Username, *foundUser.UserType, foundUser.Id)
		helper.UpdateAllTokens(token, refreshToken, foundUser.Id)
		findAfterUpdateTokensError := userCollection.FindOne(ctx, bson.M{"_id": foundUser.Id}).Decode(&foundUser)

		if findAfterUpdateTokensError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": findAfterUpdateTokensError.Error()})
			return
		}

		c.JSON(http.StatusOK, foundUser)
	}
}

// Get Users godoc
// @Sumary Get Users
// @Description Returns all users
// @Param        recordPerPage    query     int  false  "number of results per page"
// @Param        page    query     int  false  "page number"
// @Produce application/json
// @Tags users
// @Success 200
// @Failure 400
// @Failure 500
// @Router /users [get]
// @Security BearerAuth
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.CheckUserType(c, "ADMIN")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))

		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err2 := strconv.Atoi(c.Query("page"))

		if err2 != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage

		matchStage := bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{
			{"_id", bson.D{{"_id", "null"}}},
			{"total_count", bson.D{{"$sum", 1}}},
			{"data", bson.D{{"$push", "$$ROOT"}}}}}}
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}}}}}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		}

		var allusers []bson.M

		err = result.All(ctx, &allusers)

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allusers[0])
	}
}

// GetUser godoc
// @Sumary Get User
// @Description Returns object with requested id
// @Param        id   path      string  true  "User Id"
// @Produce application/json
// @Tags users
// @Success 200
// @Failure 400
// @Failure 500
// @Router /users/{id} [get]
// @Security BearerAuth
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := helper.MatchUserTypeToId(c, id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		objId, convertError := primitive.ObjectIDFromHex(id)

		if convertError != nil {
			panic(err)
		}

		findError := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

		defer cancel()

		if findError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": findError.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
