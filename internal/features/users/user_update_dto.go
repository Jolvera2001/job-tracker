package users

type UserUpdateDto struct {
	Username string `json:"username" bson:"username"`
}
