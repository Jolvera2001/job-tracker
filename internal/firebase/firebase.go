package firebase

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var Firebase_App *firebase.App
var Auth_Client *auth.Client

func InitFirebase() error {
	var firebaseApp *firebase.App
	var options option.ClientOption
	var ctx = context.Background()

	environment := os.Getenv("GO_ENV")

	if environment == "release" {
		jsonBytes, err := base64.StdEncoding.DecodeString(os.Getenv("FIREBASE_B64"))
		if err != nil {
			log.Fatalln("Error decoding base64 string!")
		}
		options = option.WithCredentialsJSON(jsonBytes)

		firebaseApp, err = firebase.NewApp(ctx, nil, options)
		if err != nil {
			log.Fatalln("Error starting Firebase app")
			return err
		}
	} else {
		var err error
		creds := os.Getenv("AUTH_SECRET")
		options = option.WithCredentialsFile(creds)

		firebaseApp, err = firebase.NewApp(ctx, nil, options)
		if err != nil {
			log.Fatalln("Error starting Firebase app")
			return err
		}
	}

	authClient, err := firebaseApp.Auth(ctx)
	if err != nil {
		log.Fatalln("Error starting Auth Client")
		return err
	}

	Firebase_App = firebaseApp
	Auth_Client = authClient

	return nil
}
