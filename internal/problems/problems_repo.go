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

func (r *Repo) GetAll(ctx context.Context) (*[]models.Problem, error) {
	var problems []models.Problem
	if err := r.DB.WithContext(ctx).Find(&problems).Error; err != nil {
		return nil, err
	}
	return &problems, nil
}

func (r *Repo) GetByID(ctx context.Context, id string) (*models.Problem, error) {
	var problems models.Problem
	if err := r.DB.WithContext(ctx).Preload("Comments").Preload("Msme").First(&problems, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &problems, nil
}
