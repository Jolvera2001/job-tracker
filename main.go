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
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting server...")

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not present...")
	}

	environment := os.Getenv("GO_ENV")
	var address string

	if environment == "release" {
		address = "0.0.0.0:8080"
		gin.SetMode(gin.ReleaseMode)
	} else {
		address = ":8080"
		gin.SetMode(gin.DebugMode)
	}


	// Initializing outide connections
	go func() {
		if err := database.ConnectToMongoDB(); err != nil {
			log.Fatalln("MongoDB Connection failed: ", err.Error())
			return
		}
	
		if err := firebase.InitFirebase(); err != nil {
			log.Fatalln("Firebase app init failed: ", err.Error())
			return
		}
	}()

	// Router
	router := gin.New()
	router.SetTrustedProxies([]string{})

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
	log.Println("Running server on: ", address)

	if err := router.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
