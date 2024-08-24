package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func InitFirebase() (*firebase.App, *auth.Client, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading dotenv")
		return nil, nil, err
	}

	creds := os.Getenv("AUTH_SECRET")
	options := option.WithCredentialsFile(creds)

	app, err := firebase.NewApp(context.Background(), nil, options)
	if err != nil {
		log.Fatalln("Error starting Firebase app")
		return nil, nil, err
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalln("Error starting Auth Client")
		return nil, nil, err
	}

	log.Println("Established Auth Client with Firebase!")

	return app, authClient, nil
}
