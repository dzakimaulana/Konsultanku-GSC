package msme

import (
	"context"
	"konsultanku-v2/internal/firebase/auth"
	"konsultanku-v2/pkg/databases"
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

func (s *Svc) AddProfile(ctx context.Context, req AddReq, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	msme := &models.MSME{
		ID:      id,
		Name:    req.Name,
		Since:   req.Since,
		Type:    req.Type,
		Created: time.Now().Unix(),
	}

	_, err := s.MsmeRepo.AddProfile(ctx, *msme)
	if err != nil {
		return err
	}
	return nil
}

func (s *Svc) GetOwnProfile(ctx context.Context, id string) (*models.MSMEOwnResp, error) {
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

	var collab []models.CollabMsmeResp
	for _, coll := range *msme.Collaboration {
		student, err := databases.AuthMd.GetUser(ctx, coll.StudentID)
		if err != nil {
			break
		}
		std := &models.UserResponse{
			UID:         student.UID,
			Email:       student.Email,
			DisplayName: student.DisplayName,
			PhoneNumber: student.PhoneNumber,
			PhotoURL:    student.PhotoURL,
		}
		collab = append(collab, models.CollabMsmeResp{
			Student:         *std,
			InCollaboration: coll.InCollaboration,
		})
	}

	resp := &models.MSMEOwnResp{
		User:          *userResp,
		Name:          msme.Name,
		Since:         msme.Since,
		Tags:          *msme.Tags,
		Problem:       *msme.Problem,
		Collaboration: collab,
	}

	return resp, nil
}

func (s *Svc) AddedCollab(ctx context.Context, studentId string, msmeId string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	collab := &models.Collaboration{
		MsmeID:          msmeId,
		StudentID:       studentId,
		InCollaboration: false,
		Finished:        false,
	}
	_, err := s.MsmeRepo.AddedCollab(ctx, *collab)
	if err != nil {
		return err
	}
	return nil
}

func (s *Svc) GiveProgress(ctx context.Context, req UpdateProgress, msmeId string, studentId string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	collab := &models.Collaboration{
		MsmeID:      msmeId,
		StudentID:   studentId,
		Progress:    int8(req.Progress % 100),
		Description: req.Description,
	}
	_, err := s.MsmeRepo.GiveProgress(ctx, *collab)
	if err != nil {
		return err
	}
	return nil
}

func (s *Svc) EndCollaboration(ctx context.Context, req EndCollaboration, msmeId string, studentId string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	collab := &models.Collaboration{
		MsmeID:      msmeId,
		StudentID:   studentId,
		Progress:    int8(100),
		Description: req.Description,
		Feedback:    req.Feedback,
		Rating:      float32(req.Rating % 5),
		Finished:    true,
	}

	var wg sync.WaitGroup
	var errStudent, errProgress error

	wg.Add(2)

	go func() {
		defer wg.Done()
		student, err := s.MsmeRepo.GetStudent(ctx, studentId)
		if err == nil {
			student = &models.Student{
				ID:     studentId,
				Rating: student.Rating + float32(req.Rating%5),
			}
			_, errStudent = s.MsmeRepo.StudentRating(ctx, *student)
		}
		errStudent = err
	}()

	go func() {
		defer wg.Done()
		_, errProgress = s.MsmeRepo.GiveProgress(ctx, *collab)
	}()

	wg.Wait()

	if errStudent != nil {
		return errStudent
	}
	if errProgress != nil {
		return errProgress
	}
	return nil
}
