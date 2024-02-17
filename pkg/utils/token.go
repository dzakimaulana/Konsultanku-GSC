package utils

import (
	"context"
	"konsultanku-v2/pkg/databases"
)

const (
	apiKey = "AIzaSyDS_P0cgEBcinAaHhb5d-vCgFhLe4AP9MU"
)

func VerifyToken(ctx context.Context, token string) error {
	_, err := databases.AuthMd.VerifyIDToken(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func RefreshToken(refreshToken string) (map[string]interface{}, error) {
	url := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=" + apiKey
	jsonData := map[string]interface{}{
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}

	resp, err := HitAPI("POST", url, jsonData)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
