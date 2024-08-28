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
	verifiedToken, err := verifyToken(c)
	if err != nil {
		return UserModel{}, fmt.Errorf("Issue with token verification")
	}

	filter := bson.M{"user_id": verifiedToken.UID}
	var userToReturn UserModel
	err = database.GetCollection("Users").FindOne(context.Background(), filter).Decode(&userToReturn)
	if err != nil {
		return UserModel{}, err
	}

	return userToReturn, nil
}

func UpdateUserService(c *gin.Context, update UserUpdateDto) (UserModel, error) {
	verifiedToken, err := verifyToken(c)
	if err != nil {
		return UserModel{}, fmt.Errorf("Issue with token verification")
	}

	var result UserModel
	filter := bson.M{"user_id": verifiedToken.UID}
	updateData, err := bson.Marshal(update)
	if err != nil {
		return UserModel{}, err
	}

	var bsonUpdate bson.M
	if err = bson.Unmarshal(updateData, &bsonUpdate); err != nil {
		return UserModel{}, err
	}

	err = database.GetCollection("Users").FindOneAndUpdate(context.Background(), filter, bsonUpdate).Decode(&result)
	if err != nil {
		return UserModel{}, err
	}

	return result, nil
}

func DeleteUserService(c *gin.Context) error {

}

func verifyToken(c *gin.Context) (*auth.Token, error) {
	user, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("Token does not exist in current context")
	}

	token, ok := user.(string)
	if !ok {
		return nil, fmt.Errorf("Current token is not compatible")
	}

	verifiedToken, err := firebase.Auth_Client.VerifyIDToken(context.Background(), token)
	if err != nil {
		return &auth.Token{}, err
	}

	return verifiedToken, err
}
