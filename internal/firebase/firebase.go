package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var Firebase_App *firebase.App
var Auth_Client *auth.Client

func InitFirebase() error {
	creds := os.Getenv("AUTH_SECRET")
	options := option.WithCredentialsFile(creds)

	app, err := firebase.NewApp(context.Background(), nil, options)
	if err != nil {
		log.Fatalln("Error starting Firebase app")
		return err
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalln("Error starting Auth Client")
		return err
	}

	log.Println("Established Auth Client with Firebase!")

	Firebase_App = app
	Auth_Client = authClient

	return nil
}
