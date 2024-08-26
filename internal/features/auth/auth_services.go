package auth

import (
	"firebase.google.com/go/auth"
)

func RegisterService(registerDto UserRegisterDto) (string, error) {
	params := (&auth.UserToCreate{}).
		Email(registerDto.Email).
		Password(registerDto.Password)

	// create user record
	// create custom token using user record uid
	// create user within mongoDB and store if and only if the previous two are successful

	// return token, user object ID, and nil error
}

func LoginService(loginDto UserLoginDto) (string, error) {

}
