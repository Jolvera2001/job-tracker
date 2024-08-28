package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"job-tracker/internal/database"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterService(registerDto UserRegisterDto) (*FirebaseApiResponse, error) {
	registerRequest := FirebaseApiRequest{
		Email:             registerDto.Email,
		Password:          registerDto.Password,
		ReturnSecureToken: true,
	}

	response, err := registerWithFirebase(registerRequest)
	if err != nil {
		return nil, err
	}

	userId := response.UID
	newUser := UserModel{
		ID:          primitive.NewObjectID(),
		UserID:      userId,
		Username:    registerDto.Username,
		Email:       registerDto.Email,
		DateCreated: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err = database.GetCollection("Users").InsertOne(context.Background(), newUser)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func LoginService(loginDto UserLoginDto) (*FirebaseApiResponse, error) {
	loginRequest := FirebaseApiRequest{
		Email:             loginDto.Email,
		Password:          loginDto.Password,
		ReturnSecureToken: true,
	}

	response, err := loginWithFirebase(loginRequest)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func registerWithFirebase(request FirebaseApiRequest) (*FirebaseApiResponse, error) {
	creds := os.Getenv("AUTH_API_KEY")
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
		var errorResponse map[string]interface{}
		json.NewDecoder(response.Body).Decode(&errorResponse)
		return nil, fmt.Errorf("status not OK: %s", errorResponse)
	}

	var registerResponse FirebaseApiResponse
	err = json.NewDecoder(response.Body).Decode(&registerResponse)
	if err != nil {
		return nil, err
	}

	return &registerResponse, nil
}

func loginWithFirebase(request FirebaseApiRequest) (*FirebaseApiResponse, error) {
	creds := os.Getenv("AUTH_API_KEY")
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
		var errorResponse map[string]interface{}
		json.NewDecoder(response.Body).Decode(&errorResponse)
		return nil, fmt.Errorf("status not OK: %s", errorResponse)
	}

	var loginResponse FirebaseApiResponse
	err = json.NewDecoder(response.Body).Decode(&loginResponse)
	if err != nil {
		return nil, err
	}

	return &loginResponse, nil
}
