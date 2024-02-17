package student

import (
	"context"
	"konsultanku-v2/pkg/models"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) StudentRepo {
	return &Repo{
		DB: db,
	}
}

func (r *Repo) GetByID(ctx context.Context, id string) (*models.Student, error) {
	var student models.Student
	if err := r.DB.WithContext(ctx).Preload("Tags").Preload("Team").First(&student, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *Repo) GetAll(ctx context.Context) (*[]models.Student, error) {
	var students []models.Student
	if err := r.DB.WithContext(ctx).Find(&students).Error; err != nil {
		return nil, err
	}
	return &students, nil
}
