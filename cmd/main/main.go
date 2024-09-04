package main

import (
	"job-tracker/internal/database"
	"job-tracker/internal/firebase"
	"log"

	"job-tracker/internal/features/applications"
	"job-tracker/internal/features/auth"
	"job-tracker/internal/features/batches"
	"job-tracker/internal/features/users"

	"github.com/gin-gonic/gin"
)

var Environment string

func main() {
	log.Println("Starting server...")

	// Initializing outide connections
	if err := database.ConnectToMongoDB(); err != nil {
		log.Fatalln("MongoDB Connection failed: ", err.Error())
		return
	}

	if err := firebase.InitFirebase(); err != nil {
		log.Fatalln("Firebase app init failed: ", err.Error())
		return
	}

	// Router
	router := gin.New()

	// add default middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

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

	log.Println("Setup successful")
	log.Println("Running server on: 0.0.0.0:8080")

	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
