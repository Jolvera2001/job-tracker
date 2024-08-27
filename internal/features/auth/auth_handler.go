package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthHandlers(r *gin.Engine) {
	v1 := r.Group("api/v1/auth")
	{
		v1.POST("/register", RegisterHandler)
		v1.POST("/login", LoginHandler)
	}
}

func RegisterHandler(c *gin.Context) {
	var registerDto UserRegisterDto
	if err := c.BindJSON(&registerDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})

	}

	response, id, err := RegisterService(registerDto)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"Auth response": response, "Id": id})
}

func LoginHandler(c *gin.Context) {
	var loginDto UserLoginDto
	if err := c.BindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	response, id, err := LoginService(loginDto)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"Auth response": response, "Id": id})
}
