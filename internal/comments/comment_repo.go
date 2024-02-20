package comments

import (
	"context"
	"konsultanku-v2/pkg/models"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		DB: db,
	}
}

func (r *Repo) AddComment(ctx context.Context, comment models.Comment) (*models.Comment, error) {
	if err := r.DB.WithContext(ctx).Create(comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *Repo) GetComments(ctx context.Context, mseId string) (*[]models.Comment, error) {
	var comments []models.Comment
	if err := r.DB.WithContext(ctx).Preload("Team").Preload("Problem").Preload("Student").Find(comments, "mse_id = ?", mseId).Error; err != nil {
		return nil, err
	}
	return &comments, nil
}
