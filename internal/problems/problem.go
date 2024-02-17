package problems

import (
	"context"
	"konsultanku-v2/pkg/models"
)

type ProblemRepo interface {
	GetAll(ctx context.Context) (*[]models.Problem, error)
	GetByID(ctx context.Context, id string) (*models.Problem, error)
}

type ProblemSvc interface {
}
