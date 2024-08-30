package applications

import (
	"job-tracker/internal/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GroupApplicationHandlers(r *gin.Engine) {
	v1 := r.Group("api/v1", middleware.AuthMiddleware())
	{
		v1.GET("/application/:appId", GetAppHandler)
		v1.GET("/application/:batchId", GetAppAllHandler)
		v1.POST("/application", CreateAppHandler)
		v1.PUT("/application", UpdateAppHandler)
		v1.DELETE("/application/:appId", DeleteAppHandler)
	}
}

func GetAppHandler(c *gin.Context) {
	appId := c.Param("appId")

	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "issue with Id"})
		return
	}

	res, err := GetAppService(id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal issue"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"application": res})
}

func GetAppAllHandler(c *gin.Context) {
	batchId := c.Param("batchId")

	id, err := primitive.ObjectIDFromHex(batchId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "issue with Id"})
		return
	}

	res, err := GetAppAllService(id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal issue"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"applications": res})
}

func CreateAppHandler(c *gin.Context) {
	var appDto AppDto
	if err := c.BindJSON(&appDto); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameters"})
		return
	}

	res, err := CreateAppService(appDto)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal issue"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"application": res})
}

func UpdateAppHandler(c *gin.Context) {
	var update AppModel
	if err := c.BindJSON(&update); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameters"})
		return
	}

	res, err := UpdateAppService(update)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal issue"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"application": res})
}

func DeleteAppHandler(c *gin.Context) {
	appId := c.Param("appId")

	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameters"})
		return
	}

	err = DeleteAppService(id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal issue"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
