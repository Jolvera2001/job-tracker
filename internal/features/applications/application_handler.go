package applications

import (
	"job-tracker/internal/middleware"

	"github.com/gin-gonic/gin"
)

func GroupApplicationHandlers(r *gin.Engine) {
	v1 := r.Group("api/v1", middleware.AuthMiddleware())
	{
		v1.GET("/application/:appId", GetAppHandler)
		v1.GET("/application", GetAppAllHandler)
		v1.POST("/application", CreateAppHandler)
		v1.PUT("/application", UpdateAppHandler)
		v1.DELETE("/application/:appId", DeleteAppHandler)
	}
}

func GetAppHandler(c *gin.Context) {

}

func GetAppAllHandler(c *gin.Context) {
	
}

func CreateAppHandler(c *gin.Context) {
	
}

func UpdateAppHandler(c *gin.Context) {
	
}

func DeleteAppHandler(c *gin.Context) {
	
}