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
	token, err := extractToken(c)
	if err != nil {
		return UserModel{}, fmt.Errorf("issue with token verification")
	}

	filter := bson.M{"user_id": token.UID}
	var userToReturn UserModel
	err = database.GetCollection("Users").FindOne(context.Background(), filter).Decode(&userToReturn)
	if err != nil {
		return UserModel{}, err
	}

	return userToReturn, nil
}

func UpdateUserService(c *gin.Context, update UserUpdateDto) (UserModel, error) {
	token, err := extractToken(c)
	if err != nil {
		return UserModel{}, fmt.Errorf("issue with token verification")
	}

	filter := bson.M{"user_id": token.UID}
	updateData, err := bson.Marshal(update)
	if err != nil {
		return UserModel{}, err
	}

	var result UserModel
	var bsonUpdate bson.M
	if err = bson.Unmarshal(updateData, &bsonUpdate); err != nil {
		return UserModel{}, err
	}

	_, err = database.GetCollection("Users").UpdateOne(context.Background(), filter, bson.M{"$set": bsonUpdate})
	if err != nil {
		return UserModel{}, err
	}

	err = database.GetCollection("Users").FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return UserModel{}, err
	}

	return result, nil
}

func DeleteUserService(c *gin.Context) error {
	token, err := extractToken(c)
	if err != nil {
		return err
	}

	filter := bson.M{"user_id": token.UID}
	_, err = database.GetCollection("Users").DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	err = firebase.Auth_Client.DeleteUser(context.Background(), token.UID)
	if err != nil {
		return err
	}

	return nil
}

func extractToken(c *gin.Context) (*auth.Token, error) {
	user, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("token does not exist in current context")
	}

	token, ok := user.(*auth.Token)
	if !ok {
		return nil, fmt.Errorf("current token is not compatible")
	}

	return token, nil
}
