package applications

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppDto struct {
	BatchId     primitive.ObjectID `json:"batch_id"`
	Name        string             `json:"name"`
	Company     string             `json:"company"`
	Description string             `json:"description"`
	Status      string             `json:"status"`
	RoundCount  int                `json:"round_count"`
}
