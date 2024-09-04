package main

import (
	"job-tracker/internal/database"
	"job-tracker/internal/firebase"
	"log"
	"os"

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
		log.Fatalln("MongoDB Connection failed")
		return
	}

	if err := firebase.InitFirebase(); err != nil {
		log.Fatalln("Firebase app init failed")
		return
	}

	// checking environment
	environment := os.Getenv("GO_ENV")

	if environment == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Router
	router := gin.New()

	// add default middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// trusted proxies
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

	log.Println("Setup successful")
	
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
