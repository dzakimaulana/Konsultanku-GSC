package student

import (
	"context"
	"konsultanku-v2/internal/firebase/auth"
	"konsultanku-v2/pkg/models"
	"sync"
	"time"
)

type Svc struct {
	StudentRepo
	AuthRepo auth.AuthRepo
	timeout  time.Duration
}

func NewSvc(sr StudentRepo, ar auth.AuthRepo) StudentSvc {
	return &Svc{
		StudentRepo: sr,
		AuthRepo:    ar,
		timeout:     time.Duration(10) * time.Second,
	}
}

func (s *Svc) GetByID(ctx context.Context, id string) (*models.StudentResponse, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var user *models.Person
	var student *models.Student
	var userErr, studentErr error

	go func() {
		defer wg.Done()
		user, userErr = s.AuthRepo.GetUserInfo(ctx, id)
	}()

	go func() {
		defer wg.Done()
		student, studentErr = s.StudentRepo.GetByID(ctx, id)
	}()

	wg.Wait()

	if userErr != nil {
		return nil, userErr
	}
	if studentErr != nil {
		return nil, studentErr
	}

	userResp := &models.UserResponse{
		UID:         user.UID,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		PhoneNumber: user.PhoneNumber,
		PhotoURL:    user.PhotoURL,
	}

	resp := &models.StudentResponse{
		User:       *userResp,
		Major:      student.Major,
		ClassOf:    student.ClassOf,
		University: student.University,
		Tags:       *student.Tags,
		Team:       *student.Team,
	}

	return resp, nil
}
