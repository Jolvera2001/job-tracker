package auth

import (
	"firebase.google.com/go/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterService(registerDto UserRegisterDto) (string, primitive.ObjectID, error) {
	params := (&auth.UserToCreate{}).
		Email(registerDto.Email).
		Password(registerDto.Password)

	// create user record
	// create custom token using user record uid
	// create user within mongoDB and store if and only if the previous two are successful

	// return token, user object ID, and nil error
}

func LoginService(loginDto UserLoginDto) (string, primitive.ObjectID, error) {
	// check for user record
	// create custom token
	// check for user object Id

	// return token, user Object ID, and nil error
}
