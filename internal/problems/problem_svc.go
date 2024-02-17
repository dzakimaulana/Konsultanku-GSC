package problems

import (
	"context"
	"konsultanku-v2/pkg/models"
	"time"
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

func (s *Svc) GetAll(ctx context.Context) (*[]models.ProblemAllRes, error) {
	fromSvc, err := s.ProblemRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var resp []models.ProblemAllRes
	for _, data := range *fromSvc {
		resp = append(resp, models.ProblemAllRes{
			ID:      data.ID,
			Like:    data.Like,
			Created: data.Created,
			Content: data.Content,
			Msme:    *data.Msme,
		})
	}
	return &resp, nil
}

func (s *Svc) GetByID(ctx context.Context, id string) (*models.ProblemIDRes, error) {
	fromSvc, err := s.ProblemRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := &models.ProblemIDRes{
		ID:       fromSvc.ID,
		Like:     fromSvc.Like,
		Created:  fromSvc.Created,
		Content:  fromSvc.Content,
		Comments: *fromSvc.Comments,
		Msme:     *fromSvc.Msme,
	}
	return resp, nil
}
