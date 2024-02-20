package problems

import (
	"context"
	"konsultanku-v2/pkg/databases"
	"konsultanku-v2/pkg/models"
	"sync"
	"time"

	"firebase.google.com/go/auth"
	"github.com/google/uuid"
)

type Svc struct {
	ProblemRepo
	timeout time.Duration
}

func NewSvc(pr ProblemRepo) ProblemSvc {
	return &Svc{
		ProblemRepo: pr,
		timeout:     time.Duration(10) * time.Second,
	}
}

func (s *Svc) GetProblem(ctx context.Context, id string) (*[]models.ProblemAllResp, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	allProblem, err := s.ProblemRepo.GetAll(ctx, id)
	if err != nil {
		return nil, err
	}

	var resp []models.ProblemAllResp
	var wg sync.WaitGroup

	for _, problem := range *allProblem {

		wg.Add(2)

		var person *auth.UserRecord
		var msme *models.MSME
		var errPerson, errMsme error

		go func() {
			defer wg.Done()
			person, errPerson = databases.AuthMd.GetUser(ctx, problem.MsmeID)
		}()

		go func() {
			defer wg.Done()
			msme, errMsme = s.ProblemRepo.GetMsmeById(ctx, problem.MsmeID)
		}()

		wg.Wait()

		if errPerson != nil || errMsme != nil {
			break
		}

		userResp := &models.UserResponse{
			UID:         person.UID,
			Email:       person.Email,
			DisplayName: person.DisplayName,
			PhoneNumber: person.PhoneNumber,
			PhotoURL:    person.PhotoURL,
		}

		msmeResp := &models.MSMEProbResp{
			User:  *userResp,
			Name:  msme.Name,
			Since: msme.Since,
		}

		resp = append(resp, models.ProblemAllResp{
			ID:           problem.ID,
			Like:         problem.Like,
			CommentCount: int64(len(*problem.Comments)),
			Created:      problem.Created,
			Title:        problem.Title,
			Content:      problem.Content,
			Msme:         *msmeResp,
		})
	}
	return &resp, nil
}

func (s *Svc) AddProblem(ctx context.Context, prob AddProblem, mseId string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	// define UID
	problemId := uuid.New()

	var wg sync.WaitGroup
	var errProb, errProbTag error

	wg.Add(2)

	go func() {
		defer wg.Done()
		problem := &models.Problem{
			ID:      problemId,
			MsmeID:  mseId,
			Like:    0,
			Created: time.Now().Unix(),
			Title:   prob.Title,
			Content: prob.Content,
			Active:  true,
		}
		_, errProb = s.ProblemRepo.AddProblem(ctx, *problem)
	}()

	go func() {
		defer wg.Done()
		var probTag []models.ProblemsTags
		for _, tagId := range prob.Tag {
			probTag = append(probTag, models.ProblemsTags{
				ProblemID: problemId.String(),
				TagID:     int64(tagId),
			})
		}
		_, errProbTag = s.ProblemRepo.AddProblemsTags(ctx, probTag)
	}()

	wg.Wait()

	if errProb != nil {
		return errProb
	}
	if errProbTag != nil {
		return errProbTag
	}
	return nil
}
