package applications

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppModel struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	BatchId     primitive.ObjectID `json:"batch_id" bson:"batch_id"`
	Name        string             `json:"name" bson:"name"`
	Company     string             `json:"company" bson:"company"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	RoundCount  int                `json:"round_count" bson:"round_count"`
}
