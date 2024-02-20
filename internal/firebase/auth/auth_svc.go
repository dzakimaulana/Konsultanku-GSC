package auth

import (
	"context"
	"errors"
	"konsultanku-v2/internal/firebase/storage"
	"konsultanku-v2/pkg/models"
	"time"
)

type Svc struct {
	AuthRepo
	StorageRepo storage.StorageRepo
	timeout     time.Duration
}

func NewSvc(ra AuthRepo, rs storage.StorageRepo) AuthSvc {
	return &Svc{
		AuthRepo:    ra,
		StorageRepo: rs,
		timeout:     time.Duration(10) * time.Second,
	}
}

func (s *Svc) Register(ctx context.Context, regis RegisterFirebaseReq) (*models.EmailRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	// convert into Person struct
	user := &models.Person{
		Email:         regis.Email,
		EmailVerified: false,
		DisplayName:   regis.DisplayName,
		Password:      regis.Password,
		PhoneNumber:   regis.PhoneNumber,
		PhotoURL:      regis.PhotoURL,
		Disabled:      false,
		CustomClaims: map[string]interface{}{
			"role": regis.Role,
		},
	}

	// Get Token and register user
	resp, err := s.AuthRepo.Register(ctx, *user)
	if err != nil {
		return nil, err
	}

	// Upload file into Firebase Storage
	fileURL, err := s.StorageRepo.UploadFile(ctx, *regis.File)
	if err != nil {
		return nil, err
	}
	user.PhotoURL = fileURL // Get file URL

	// Update data into Firebase Auth
	user.UID = resp["localId"].(string)
	_, err = s.AuthRepo.UpdateAuthData(ctx, *user)
	if err != nil {
		return nil, err
	}

	// Send email verification
	token := resp["idToken"].(string)
	evResp, err := s.AuthRepo.EmailVerification(ctx, token)
	if err != nil {
		return nil, err
	}

	finalResponse := &models.EmailRes{
		Email: evResp["email"].(string),
	}

	// return token for email verification
	return finalResponse, nil
}

func (s *Svc) Login(ctx context.Context, log LoginReq) (*models.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	// convert into Person struct
	user := &models.Person{
		Email:    log.Email,
		Password: log.Password,
	}

	// login from google api
	resp, err := s.AuthRepo.Login(ctx, *user)
	if err != nil {
		return nil, err
	}

	// check email verified
	uid := resp["localId"].(string)
	usr, err := s.AuthRepo.GetUserInfo(ctx, uid)
	if err != nil {
		return nil, err
	}
	if usr.EmailVerified == false {
		return nil, errors.New("Verify your email first!!")
	}

	userData := &models.UserResponse{
		UID:         usr.UID,
		Email:       usr.Email,
		DisplayName: usr.DisplayName,
		PhoneNumber: usr.PhoneNumber,
		PhotoURL:    usr.PhotoURL,
	}

	// give user token
	logResp := &models.AuthResponse{
		IDToken:      resp["idToken"].(string),
		RefreshToken: resp["refreshToken"].(string),
		LocalID:      resp["localId"].(string),
		User:         userData,
	}
	return logResp, nil
}

func (s *Svc) ResetPassword(ctx context.Context, rp RPReq) (*models.EmailRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	resp, err := s.AuthRepo.ResetPassword(ctx, rp.Email)
	if err != nil {
		return nil, err
	}

	er := &models.EmailRes{
		Email: resp["email"].(string),
	}
	return er, nil
}

func (s *Svc) Logout(ctx context.Context, uid string) error {
	if err := s.AuthRepo.Logout(ctx, uid); err != nil {
		return err
	}
	return nil
}
