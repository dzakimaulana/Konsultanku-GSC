package problems

import (
	"context"
	"konsultanku-v2/pkg/models"
)

type AddProblem struct {
	Title   string `json:"title"`
	Tag     []int  `json:"tag"`
	Content string `json:"content"`
}

type ProblemRepo interface {
	GetAll(ctx context.Context, id string) (*[]models.Problem, error)
	AddProblem(ctx context.Context, prob models.Problem) (*models.Problem, error)
	AddProblemsTags(ctx context.Context, pt []models.ProblemsTags) (*[]models.ProblemsTags, error)
	GetMsmeById(ctx context.Context, id string) (*models.MSME, error)
}

type ProblemSvc interface {
	AddProblem(ctx context.Context, prob AddProblem, mseId string) error
	GetProblem(ctx context.Context, id string) (*[]models.ProblemAllResp, error)
}
