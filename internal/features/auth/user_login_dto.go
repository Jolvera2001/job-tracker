package auth

type UserLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
