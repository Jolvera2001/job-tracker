package batches

import (
	"job-tracker/internal/middleware"

	"github.com/gin-gonic/gin"
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

}

func GetBatchAllHandler(c *gin.Context) {
	
}

func CreateBatchHandler(c *gin.Context) {
	
}

func UpdateBatchHandler(c *gin.Context) {
	
}

func DeleteBatchHandler(c *gin.Context) {
	
}