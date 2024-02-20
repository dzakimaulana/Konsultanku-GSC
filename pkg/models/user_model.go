package models

type Person struct {
	UID           string
	Email         string
	EmailVerified bool
	DisplayName   string
	Password      string
	PhoneNumber   string
	PhotoURL      string
	Disabled      bool
	CustomClaims  map[string]interface{}
}

type UserResponse struct {
	UID         string `json:"uid"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	PhoneNumber string `json:"phone_number"`
	PhotoURL    string `json:"photo_url"`
}

type UserPhoneResp struct {
	PhoneNumber string `json:"phone_number"`
}

type AuthResponse struct {
	IDToken      string        `json:"id_token"`
	RefreshToken string        `json:"refresh_token"`
	LocalID      string        `json:"local_id"`
	User         *UserResponse `json:"user"`
}

type EmailRes struct {
	Email string `json:"email"`
}
