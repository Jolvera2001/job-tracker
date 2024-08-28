package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserUpdateDto struct {
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	DateCreated primitive.DateTime `json:"date_created" bson:"date_created"`
}