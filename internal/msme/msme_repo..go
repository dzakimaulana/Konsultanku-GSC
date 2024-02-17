package msme

import (
	"context"
	"konsultanku-v2/pkg/models"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) MsmeRepo {
	return &Repo{
		DB: db,
	}
}

func (r *Repo) AddProfile(ctx context.Context, msme models.MSME) (*models.MSME, error) {
	if err := r.DB.WithContext(ctx).Create(msme).Error; err != nil {
		return nil, err
	}
	return &msme, nil
}

func (r *Repo) UpdateProfile(ctx context.Context, msme models.MSME) (*models.MSME, error) {
	if err := r.DB.WithContext(ctx).Updates(msme).Error; err != nil {
		return nil, err
	}
	return &msme, nil
}

func (r *Repo) GetByID(ctx context.Context, id string) (*models.MSME, error) {
	var msme models.MSME
	if err := r.DB.WithContext(ctx).Preload("Tags").Preload("Problem").First(&msme, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &msme, nil
}

func (r *Repo) GetAll(ctx context.Context) (*[]models.MSME, error) {
	var msmes []models.MSME
	if err := r.DB.WithContext(ctx).Find(&msmes).Error; err != nil {
		return nil, err
	}
	return &msmes, nil
}
