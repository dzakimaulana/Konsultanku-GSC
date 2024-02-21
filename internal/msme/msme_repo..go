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
	if err := r.DB.WithContext(ctx).Preload("Tags").Preload("Problem").Preload("Collaboration").First(&msme, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &msme, nil
}

func (r *Repo) AddedCollab(ctx context.Context, cl models.Collaboration) (*models.Collaboration, error) {
	if err := r.DB.WithContext(ctx).Create(cl).Error; err != nil {
		return nil, err
	}
	return &cl, nil
}

func (r *Repo) GiveProgress(ctx context.Context, collab models.Collaboration) (*models.Collaboration, error) {
	if err := r.DB.WithContext(ctx).Save(collab).Error; err != nil {
		return nil, err
	}
	return &collab, nil
}

func (r *Repo) StudentRating(ctx context.Context, student models.Student) (*models.Student, error) {
	if err := r.DB.WithContext(ctx).Save(student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *Repo) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	var student models.Student
	if err := r.DB.WithContext(ctx).First(&student, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *Repo) FindProblemTags(ctx context.Context, id string) (*[]models.ProblemsTags, error) {
	var pt []models.ProblemsTags
	if err := r.DB.WithContext(ctx).Find(&pt, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &pt, nil
}
