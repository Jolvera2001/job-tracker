package users

import (
	"job-tracker/internal/middleware"

	"github.com/gin-gonic/gin"
)

func GroupUserHandlers(r *gin.Engine) {
	v1 := r.Group("api/v1/", middleware.AuthMiddleware())
	{
		v1.GET("/user", GetUserHandler)
		v1.PUT("/user", UpdateUserHandler)
		v1.DELETE("/user", DeleteUserHandler)
	}
}

func GetUserHandler(c *gin.Context) {

}

func UpdateUserHandler(c *gin.Context) {
	
}

func DeleteUserHandler(c *gin.Context) {
	
}