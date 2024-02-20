package problems

import (
	"context"
	"konsultanku-v2/pkg/models"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) ProblemRepo {
	return &Repo{
		DB: db,
	}
}

func (r *Repo) GetAll(ctx context.Context, id string) (*[]models.Problem, error) {
	var problems []models.Problem
	if err := r.DB.WithContext(ctx).Where("id = ?", id).Find(&problems).Error; err != nil {
		return nil, err
	}
	return &problems, nil
}

func (r *Repo) AddProblem(ctx context.Context, prob models.Problem) (*models.Problem, error) {
	if err := r.DB.WithContext(ctx).Create(&prob).Error; err != nil {
		return nil, err
	}
	return &prob, nil
}

func (r *Repo) AddProblemsTags(ctx context.Context, pt []models.ProblemsTags) (*[]models.ProblemsTags, error) {
	if err := r.DB.WithContext(ctx).Create(&pt).Error; err != nil {
		return nil, err
	}
	return &pt, nil
}

func (r *Repo) GetMsmeById(ctx context.Context, id string) (*models.MSME, error) {
	var msme models.MSME
	if err := r.DB.WithContext(ctx).Preload("Tags").Preload("Comments").First(&msme, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &msme, nil
}
