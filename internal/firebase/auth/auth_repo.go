package auth

import (
	"context"
	"konsultanku-v2/pkg/models"
	"konsultanku-v2/pkg/utils"

	"firebase.google.com/go/auth"
)

const (
	apiKey = "AIzaSyDS_P0cgEBcinAaHhb5d-vCgFhLe4AP9MU"
)

type Repo struct {
	Auth *auth.Client
}

func NewRepo(a *auth.Client) AuthRepo {
	return &Repo{
		Auth: a,
	}
}

func (r *Repo) UpdateAuthData(ctx context.Context, user models.Person) (*models.Person, error) {
	params := (&auth.UserToUpdate{}).
		EmailVerified(user.EmailVerified).
		PhoneNumber(user.PhoneNumber).
		DisplayName(user.DisplayName).
		PhotoURL(user.PhotoURL).
		Disabled(user.Disabled)
	updateUser, err := r.Auth.UpdateUser(ctx, user.UID, params)
	if err != nil {
		return nil, err
	}

	if err := r.Auth.SetCustomUserClaims(ctx, updateUser.UID, map[string]interface{}{
		"role": user.CustomClaims["role"],
	}); err != nil {
		return nil, err
	}

	res := &models.Person{
		UID:           updateUser.UID,
		Email:         updateUser.Email,
		EmailVerified: updateUser.EmailVerified,
		PhoneNumber:   updateUser.PhoneNumber,
		DisplayName:   updateUser.DisplayName,
		PhotoURL:      updateUser.PhotoURL,
		Disabled:      updateUser.Disabled,
		CustomClaims:  updateUser.CustomClaims,
	}
	return res, nil
}

func (r *Repo) Register(ctx context.Context, user models.Person) (map[string]interface{}, error) {
	url := "https://identitytoolkit.googleapis.com/v1/accounts:signUp?key=" + apiKey
	jsonData := map[string]interface{}{
		"email":             user.Email,
		"password":          user.Password,
		"returnSecureToken": true,
	}

	resp, err := utils.HitAPI("POST", url, jsonData)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Repo) EmailVerification(ctx context.Context, token string) (map[string]interface{}, error) {
	url := "https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key=" + apiKey
	jsonData := map[string]interface{}{
		"requestType": "VERIFY_EMAIL",
		"idToken":     token,
	}

	resp, err := utils.HitAPI("POST", url, jsonData)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Repo) Login(ctx context.Context, user models.Person) (map[string]interface{}, error) {
	url := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=" + apiKey
	jsonData := map[string]interface{}{
		"email":             user.Email,
		"password":          user.Password,
		"returnSecureToken": true,
	}

	resp, err := utils.HitAPI("POST", url, jsonData)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Repo) GetUserInfo(ctx context.Context, uid string) (*models.Person, error) {
	resp, err := r.Auth.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	p := &models.Person{
		UID:           resp.UID,
		Email:         resp.Email,
		EmailVerified: resp.EmailVerified,
		DisplayName:   resp.DisplayName,
		PhoneNumber:   resp.PhoneNumber,
		PhotoURL:      resp.PhotoURL,
	}
	return p, nil
}

func (r *Repo) ResetPassword(ctx context.Context, email string) (map[string]interface{}, error) {
	url := "https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key=" + apiKey
	jsonData := map[string]interface{}{
		"email":       email,
		"requestType": "PASSWORD_RESET",
	}
	resp, err := utils.HitAPI("POST", url, jsonData)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Repo) Logout(ctx context.Context, uid string) error {
	if err := r.Auth.RevokeRefreshTokens(ctx, uid); err != nil {
		return err
	}
	return nil
}
