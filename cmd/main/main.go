package main

import (
	"fmt"
	"job-tracker/internal/database"
	"job-tracker/internal/firebase"
	"log"

	"job-tracker/internal/features/applications"
	"job-tracker/internal/features/auth"
	"job-tracker/internal/features/batches"
	"job-tracker/internal/features/users"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server...")

	// Initializing outide connections
	if err := database.ConnectToMongoDB(); err != nil {
		log.Fatalln("MongoDB Connection failed")
		return
	}

	if err := firebase.InitFirebase(); err != nil {
		log.Fatalln("Firebase app init failed")
		return
	}

	// Router
	router := gin.Default()
	router.SetTrustedProxies([]string{})

	// landing page
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"title": "Welcome!"})
	})

	// Handlers
	auth.GroupAuthHandlers(router)
	users.GroupUserHandlers(router)
	batches.GroupBatchHandlers(router)
	applications.GroupApplicationHandlers(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
