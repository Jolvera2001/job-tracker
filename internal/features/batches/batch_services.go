package batches

import (
	"fmt"

	"cloud.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBatchService(c *gin.Context) (BatchModel, error) {

}

func GetBatchAllService(c *gin.Context) ([]BatchModel, error) {

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
