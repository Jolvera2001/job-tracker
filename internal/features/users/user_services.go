package users

import (
	"context"
	"fmt"
	"job-tracker/internal/database"
	"job-tracker/internal/firebase"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

	filter := bson.M{"user_id": verifiedToken.UID}
	var userToReturn UserModel
	err = database.GetCollection("Users").FindOne(context.Background(), filter).Decode(&userToReturn)
	if err != nil {
		return UserModel{}, fmt.Errorf("Error finding user")
	}

	return userToReturn, nil
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
