package batches

import "go.mongodb.org/mongo-driver/bson/primitive"

type BatchModel struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	UserId      string             `json:"user_id" bson:"user_id"`
	Name        string             `json:"name" bson:"name"`
	DateCreated primitive.DateTime `json:"date_created" bson:"date_created"`
}
