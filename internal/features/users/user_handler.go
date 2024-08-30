package users

import (
	"job-tracker/internal/middleware"
	"log"
	"net/http"

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
	response, err := GetUserService(c)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": response})
}

func UpdateUserHandler(c *gin.Context) {
	var update UserUpdateDto
	if err := c.BindJSON(&update); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := UpdateUserService(c, update)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": response})
}

func DeleteUserHandler(c *gin.Context) {
	err := DeleteUserService(c)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
