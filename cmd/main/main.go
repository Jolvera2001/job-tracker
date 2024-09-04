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
	var address string

	// Initializing outide connections
	if err := database.ConnectToMongoDB(); err != nil {
		log.Fatalln("MongoDB Connection failed: ", err.Error())
		return
	}

	if err := firebase.InitFirebase(); err != nil {
		log.Fatalln("Firebase app init failed: ", err.Error())
		return
	}

	// checking environment
	environment := os.Getenv("GO_ENV")

	if environment == "release" {
		gin.SetMode(gin.ReleaseMode)
		address = "0.0.0.0:8080"
	} else {
		gin.SetMode(gin.DebugMode)
		address = ":8080"
	}

	// Router
	router := gin.New()

	// add default middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// trusted proxies
	router.SetTrustedProxies([]string{"0.0.0.0/0"})

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
	log.Println("Running server on: ", address)

	if err := router.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
