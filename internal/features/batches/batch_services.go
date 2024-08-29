package batches

import (
	"context"
	"fmt"
	"job-tracker/internal/database"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBatchService(c *gin.Context, batchId primitive.ObjectID) (BatchModel, error) {
	token, err := extractToken(c)
	if err != nil {
		return BatchModel{}, err
	}

	filter := bson.M{"user_id": token.UID, "_id": batchId}

	var batch BatchModel
	err = database.GetCollection("Batches").FindOne(context.Background(), filter).Decode(&batch)
	if err != nil {
		return BatchModel{}, err
	}

	return batch, nil
}

func GetBatchAllService(c *gin.Context) ([]BatchModel, error) {
	token, err := extractToken(c)
	if err != nil {
		return []BatchModel{}, err
	}

	filter := bson.M{"user_id": token.UID}

	var batchList []BatchModel
	cur, err := database.GetCollection("Batches").Find(context.Background(), filter)
	if err != nil {
		return []BatchModel{}, err
	}

	// look into why this is being done
	defer func() {
		if closeError := cur.Close(context.Background()); closeError != nil {
			err = closeError
		}
	}()

	err = parseCursor(&batchList, cur)
	if err != nil {
		return []BatchModel{}, err
	}

	return batchList, err
}

func CreateBatchService(c *gin.Context, newBatch BatchDto) (BatchModel, error) {

}

func UpdateBatchService(c *gin.Context, update BatchDto) (BatchModel, error) {

}

func DeleteBatchService(c *gin.Context, batchId primitive.ObjectID) (BatchModel, error) {

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

func parseCursor(batchList *[]BatchModel, cursor *mongo.Cursor) error {
	for cursor.Next(context.Background()) {
		var batch BatchModel
		if err := cursor.Decode(&batch); err != nil {
			return err
		}

		*batchList = append(*batchList, batch)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return nil
}
