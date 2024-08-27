package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"job-tracker/internal/database"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterService(registerDto UserRegisterDto) (*FirebaseApiResponse, primitive.ObjectID, error) {
	registerRequest := FirebaseApiRequest{
		Email:             registerDto.Email,
		Password:          registerDto.Password,
		ReturnSecureToken: true,
	}

	response, err := registerWithFirebase(registerRequest)
	if err != nil {
		return nil, primitive.NilObjectID, err
	}

	newUser := UserModel{
		ID:          primitive.NewObjectID(),
		Username:    registerDto.Username,
		Email:       registerDto.Email,
		DateCreated: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err = database.GetCollection("Users").InsertOne(context.Background(), newUser)
	if err != nil {
		return nil, primitive.NilObjectID, err
	}

	return response, newUser.ID, nil
}

func LoginService(loginDto UserLoginDto) (*FirebaseApiResponse, primitive.ObjectID, error) {
	loginRequest := FirebaseApiRequest{
		Email: loginDto.Email,
		Password: loginDto.Password,
		ReturnSecureToken: true,
	}

	response, err := loginWithFirebase(loginRequest)
	if err != nil {
		return nil, primitive.NilObjectID, err
	}

	var user UserModel
	var filter = bson.M{"email": loginDto.Email}
	err = database.GetCollection("Users").FindOne(context.Background(), filter).Decode(&user)

	return response, user.ID, err
}

func registerWithFirebase(request FirebaseApiRequest) (*FirebaseApiResponse, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	creds := os.Getenv("AUTH_SECRET")
	url := "https://identitytoolkit.googleapis.com/v1/accounts:signUp?key=" + creds

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Status not OK")
	}

	var registerResponse FirebaseApiResponse
	err = json.NewDecoder(response.Body).Decode(&registerResponse)
	if err != nil {
		return nil, err
	}

	return &registerResponse, nil
}

func loginWithFirebase(request FirebaseApiRequest) (*FirebaseApiResponse, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	creds := os.Getenv("AUTH_SECRET")
	url := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=" + creds

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Status not OK")
	}

	var loginResponse FirebaseApiResponse
	err = json.NewDecoder(response.Body).Decode(&loginResponse)
	if err != nil {
		return nil, err
	}

	return &loginResponse, nil
}
