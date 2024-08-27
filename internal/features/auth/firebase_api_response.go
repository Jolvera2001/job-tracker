package auth

type FirebaseApiResponse struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn`
	UID          string `json:"localId"`
}
