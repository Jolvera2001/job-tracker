package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GroupAuthHandlers(r *gin.Engine) {
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

	response, err := RegisterService(registerDto)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Auth response": response})
}

func LoginHandler(c *gin.Context) {
	var loginDto UserLoginDto
	if err := c.BindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	response, err := LoginService(loginDto)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Auth response": response})
}
