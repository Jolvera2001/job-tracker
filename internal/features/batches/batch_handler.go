package batches

import (
	"job-tracker/internal/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GroupBatchHandlers(r *gin.Engine) {
	v1 := r.Group("api/v1", middleware.AuthMiddleware())
	{
		v1.GET("/batch", GetBatchHandler)
		v1.GET("/batch/all", GetBatchAllHandler)
		v1.POST("/batch", CreateBatchHandler)
		v1.PUT("/batch", UpdateBatchHandler)
		v1.DELETE("/batch", DeleteBatchHandler)
	}
}

func GetBatchHandler(c *gin.Context) {
	var batchId primitive.ObjectID
	if err := c.BindJSON(&batchId); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Arguments"})
		return
	}

	response, err := GetBatchService(c, batchId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Batch": response})
}

func GetBatchAllHandler(c *gin.Context) {
	response, err := GetBatchAllService(c)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"list": response})
}

func CreateBatchHandler(c *gin.Context) {
	var batchDto BatchDto
	if err := c.BindJSON(&batchDto); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := CreateBatchService(c, batchDto)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"batch": response})
}

func UpdateBatchHandler(c *gin.Context) {
	
}

func DeleteBatchHandler(c *gin.Context) {
	
}