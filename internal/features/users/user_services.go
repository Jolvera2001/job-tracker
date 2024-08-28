package users

import (
	"context"
	"fmt"
	"job-tracker/internal/firebase"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func GetUserService(c *gin.Context) (UserModel, error) {
	user, exists := c.Get("user")
	if exists == false {
		return UserModel{}, fmt.Errorf("Token does not exist in current context")
	}

	token, ok := user.(string)
	if !ok {
		return UserModel{}, fmt.Errorf("Current token is not compatible")
	}

	verifiedToken, err := verifyToken(token)
	if err != nil {
		return UserModel{}, fmt.Errorf("Issue with token verification")
	}

	

	

}

func UpdateUserService() (UserModel, error) {

}

func DeleteUserService() error {

}

func verifyToken(token string) (*auth.Token, error) {
	verifiedToken, err := firebase.Auth_Client.VerifyIDToken(context.Background(), token)
	if err != nil {
		return &auth.Token{}, err
	}

	return verifiedToken, err
}
