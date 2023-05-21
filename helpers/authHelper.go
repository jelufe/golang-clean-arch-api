package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToId(c *gin.Context, id string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	err = nil

	if userType == "USER" && uid != id {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	err = CheckUserType(c, userType)
	return err
}

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")

	err = nil

	if userType != role {
		err = errors.New("Unauthorized to access this resource")
	}

	return err
}
