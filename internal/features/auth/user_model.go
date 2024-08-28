package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	UserID      string             `json:"user_id" bson:"user_id"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	DateCreated primitive.DateTime `json:"date_created" bson:"date_created"`
}
