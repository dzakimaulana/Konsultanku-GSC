package msme

import (
	"context"
	"konsultanku-v2/internal/firebase/auth"
	"konsultanku-v2/pkg/models"
	"sync"
	"time"
)

type Svc struct {
	MsmeRepo
	AuthRepo auth.AuthRepo
	timeout  time.Duration
}

func NewSvc(mr MsmeRepo, ar auth.AuthRepo) MsmeSvc {
	return &Svc{
		MsmeRepo: mr,
		AuthRepo: ar,
		timeout:  time.Duration(10) * time.Second,
	}
}

func (s *Svc) GetByID(ctx context.Context, id string) (*models.MSMEResp, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	var user *models.Person
	var msme *models.MSME
	var userErr, msmeErr error

	go func() {
		defer wg.Done()
		user, userErr = s.AuthRepo.GetUserInfo(ctx, id)
	}()

	go func() {
		defer wg.Done()
		msme, msmeErr = s.MsmeRepo.GetByID(ctx, id)
	}()

	wg.Wait()

	if userErr != nil {
		return nil, userErr
	}
	if msmeErr != nil {
		return nil, msmeErr
	}

	userResp := &models.UserResponse{
		UID:         user.UID,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		PhoneNumber: user.PhoneNumber,
		PhotoURL:    user.PhotoURL,
	}

	resp := &models.MSMEResp{
		User:    *userResp,
		Name:    msme.Name,
		Since:   msme.Since,
		Tags:    *msme.Tags,
		Problem: *msme.Problem,
	}

	return resp, nil
}

func (s *Svc) AddProfile(ctx context.Context, req AddReq, id string) (*models.MSMEResp, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	var wg sync.WaitGroup

	var user *models.Person
	var msmeResp *models.MSME
	msme := &models.MSME{
		ID:      id,
		Name:    req.Name,
		Since:   req.Since,
		Created: time.Now().Unix(),
	}
	var userErr, msmeErr error

	wg.Add(2)
	go func() {
		defer wg.Done()
		user, userErr = s.AuthRepo.GetUserInfo(ctx, id)
	}()

	go func() {
		defer wg.Done()
		msmeResp, msmeErr = s.MsmeRepo.AddProfile(ctx, *msme)
	}()
	wg.Wait()

	if userErr != nil {
		return nil, userErr
	}
	if msmeErr != nil {
		return nil, msmeErr
	}

	userResp := &models.UserResponse{
		UID:         user.UID,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		PhoneNumber: user.PhoneNumber,
		PhotoURL:    user.PhotoURL,
	}

	resp := &models.MSMEResp{
		User:  *userResp,
		Name:  msmeResp.Name,
		Since: msmeResp.Name,
	}
	return resp, nil
}
