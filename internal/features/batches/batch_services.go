package batches

import (
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
