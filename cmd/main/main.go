package main

import (
	"fmt"
	"job-tracker/internal/database"
	"job-tracker/internal/firebase"
	"log"

	"job-tracker/internal/features/auth"
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

	// Handlers
	auth.GroupAuthHandlers(router)
	users.GroupUserHandlers(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
