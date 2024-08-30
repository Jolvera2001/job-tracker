package batches

import "go.mongodb.org/mongo-driver/bson/primitive"

type BatchUpdateDto struct {
	ID   primitive.ObjectID `json:"_id"`
	Name string             `json:"name"`
}
