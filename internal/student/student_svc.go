package student

import (
	"context"
	"konsultanku-v2/internal/firebase/auth"
	"konsultanku-v2/pkg/databases"
	"konsultanku-v2/pkg/models"
	"sync"
	"time"

	"github.com/google/uuid"
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

func (s *Svc) AddProfile(ctx context.Context, studentReq AddStudent, id string) error {
	student := &models.Student{
		ID:         id,
		Major:      studentReq.Major,
		University: studentReq.University,
		ClassOf:    studentReq.ClassOf,
		Created:    time.Now().Unix(),
	}
	_, err := s.StudentRepo.AddProfile(ctx, *student)
	if err != nil {
		return err
	}
	return nil
}

func (s *Svc) GetOwnProfile(ctx context.Context, id string) (*models.StudentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

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

	var collab []models.CollabStudentResp
	for _, coll := range *student.Collaboration {
		msme, err := databases.AuthMd.GetUser(ctx, coll.MsmeID)
		if err != nil {
			break
		}
		m := &models.UserResponse{
			UID:         msme.UID,
			Email:       msme.Email,
			DisplayName: msme.DisplayName,
			PhoneNumber: msme.PhoneNumber,
			PhotoURL:    msme.PhotoURL,
		}
		collab = append(collab, models.CollabStudentResp{
			Msme:            *m,
			InCollaboration: coll.InCollaboration,
		})
	}

	resp := &models.StudentResponse{
		User:          *userResp,
		Major:         student.Major,
		ClassOf:       student.ClassOf,
		University:    student.University,
		Tags:          *student.Tags,
		Team:          *student.Team,
		Collaboration: collab,
	}

	return resp, nil
}

func (s *Svc) AcceptOffer(ctx context.Context, msmeId string, studentId string) (*models.UserPhoneResp, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	var wg sync.WaitGroup

	var msmeResp *models.UserPhoneResp
	var errUpdate, errUser error

	wg.Add(2)

	go func() {
		defer wg.Done()
		collaboration := &models.Collaboration{
			MsmeID:          msmeId,
			StudentID:       studentId,
			InCollaboration: true,
		}
		_, errUpdate = s.StudentRepo.UpdateCollaboration(ctx, *collaboration)
	}()

	go func() {
		defer wg.Done()
		msme, err := databases.AuthMd.GetUser(ctx, msmeId)
		errUser = err
		msmeResp = &models.UserPhoneResp{
			PhoneNumber: msme.PhoneNumber,
		}
	}()

	if errUpdate != nil {
		return nil, errUpdate
	}
	if errUser != nil {
		return nil, errUser
	}
	return msmeResp, nil
}

func (s *Svc) GetCollaboration(ctx context.Context, studentId string) (*[]models.GetStudentCollaboration, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	collabs, err := s.StudentRepo.GetCollaboration(ctx, studentId)
	if err != nil {
		return nil, err
	}
	var collaborations []models.GetStudentCollaboration
	for _, collab := range *collabs {
		msme, err := databases.AuthMd.GetUser(ctx, collab.MsmeID)
		if err != nil {
			break
		}
		user := &models.UserResponse{
			UID:         msme.UID,
			Email:       msme.Email,
			DisplayName: msme.DisplayName,
			PhoneNumber: msme.PhoneNumber,
			PhotoURL:    msme.PhotoURL,
		}
		collaborations = append(collaborations, models.GetStudentCollaboration{
			Msme:            *user,
			InCollaboration: collab.InCollaboration,
			Description:     collab.Description,
			Finished:        collab.Finished,
			Feedback:        collab.Feedback,
			Rating:          collab.Rating,
		})
	}
	return &collaborations, nil
}

func (s *Svc) CreateTeam(ctx context.Context, req CreateTeam) (*models.TeamResp, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	team := &models.Team{
		ID:   uuid.New(),
		Name: req.Name,
	}
	team, err := s.StudentRepo.CreateTeam(ctx, *team)
	if err != nil {
		return nil, err
	}
	teamResp := &models.TeamResp{
		ID:     team.ID.String(),
		Name:   team.Name,
		Rating: team.Rating,
	}
	return teamResp, nil
}

func (s *Svc) JoinTeam(ctx context.Context, teamId string, studentId string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	uuidTeam, err := uuid.Parse(teamId)
	if err != nil {
		return err
	}

	student := &models.Student{
		ID:     studentId,
		TeamID: uuidTeam,
	}
	_, err = s.StudentRepo.JoinTeam(ctx, *student)
	if err != nil {
		return err
	}
	return nil
}
