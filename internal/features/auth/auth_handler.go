package auth

import (
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

	// res, err := service.(registerDto) should contain token
	// if err != nil {
	// log err
	//}
	// c.JSON(http.StatusOK, gin.H{"Token": res})
}

func LoginHandler(c *gin.Context) {
	var loginDto UserLoginDto
	if err := c.BindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	// res, err := service.(loginDto) should contain token
	// if err != nil {
	// log err
	//}
	// c.JSON(http.StatusOK, gin.H{"Token": res})
}
