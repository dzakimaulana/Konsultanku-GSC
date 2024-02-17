package auth

import (
	"context"
	"konsultanku-v2/pkg/models"
	"mime/multipart"
)

type RegisterFirebaseReq struct {
	Email       string `form:"email"`
	DisplayName string `form:"name"`
	Password    string `form:"password"`
	PhoneNumber string `form:"phone_number"`
	Role        string `form:"role"`
	File        *multipart.FileHeader
	PhotoURL    string
}

type RPReq struct {
	Email string `json:"email"`
}

type UserFirebaseReq struct {
	Email       string
	DisplayName string
	Password    string
	PhoneNumber string
	Role        string
	PhotoURL    string
}

type LoginReq struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type AuthRepo interface {
	Register(ctx context.Context, user models.Person) (map[string]interface{}, error)
	UpdateAuthData(ctx context.Context, user models.Person) (*models.Person, error)
	EmailVerification(ctx context.Context, token string) (map[string]interface{}, error)
	Login(ctx context.Context, user models.Person) (map[string]interface{}, error)
	GetUserInfo(ctx context.Context, uid string) (*models.Person, error)
	ResetPassword(ctx context.Context, email string) (map[string]interface{}, error)
	Logout(ctx context.Context, uid string) error
}

type AuthSvc interface {
	Register(ctx context.Context, regis RegisterFirebaseReq) (*models.EmailRes, error)
	Login(ctx context.Context, log LoginReq) (*models.AuthResponse, error)
	ResetPassword(ctx context.Context, rp RPReq) (*models.EmailRes, error)
	Logout(ctx context.Context, uid string) error
}
