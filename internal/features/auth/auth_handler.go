package auth

import "github.com/gin-gonic/gin"

func AuthHandlers(r *gin.Engine) {
	v1 := r.Group("api/v1/auth") 
	{
		v1.POST("/register", RegisterHandler)
		v1.POST("/login", LoginHandler)
	}
}

func RegisterHandler(c *gin.Context) {

}

func LoginHandler(c *gin.Context) {

}