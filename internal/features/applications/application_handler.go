package applications

import (
	"job-tracker/internal/middleware"

	"github.com/gin-gonic/gin"
)

func GroupApplicationHandlers(r *gin.Engine) {
	v1 := r.Group("api/v1", middleware.AuthMiddleware())
	{
		v1.GET("/application/:appId", GetAppService)
		v1.GET("/application", GetAppAllService)
		v1.POST("/application", CreateAppService)
		v1.PUT("/application", UpdateAppService)
		v1.DELETE("/application/:appId", DeleteAppService)
	}
}

func GetAppService(c *gin.Context) {

}

func GetAppAllService(c *gin.Context) {
	
}

func CreateAppService(c *gin.Context) {
	
}

func UpdateAppService(c *gin.Context) {
	
}

func DeleteAppService(c *gin.Context) {
	
}