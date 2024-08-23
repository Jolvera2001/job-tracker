package main

import (
	"fmt"
	"job-tracker/internal/database"
	"job-tracker/internal/firebase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server...")

	// initializing outide connections
	database.ConnectToMongoDB()
	firebase.InitFirebase()
	router := gin.Default()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
