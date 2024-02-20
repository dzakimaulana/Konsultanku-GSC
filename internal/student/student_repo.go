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

func (r *Repo) AddProfile(ctx context.Context, student models.Student) (*models.Student, error) {
	if err := r.DB.WithContext(ctx).Create(student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *Repo) GetByID(ctx context.Context, id string) (*models.Student, error) {
	var student models.Student
	if err := r.DB.WithContext(ctx).Preload("Tags").Preload("Team").Preload("Collaboration").First(&student, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *Repo) UpdateCollaboration(ctx context.Context, collaboration models.Collaboration) (*models.Collaboration, error) {
	if err := r.DB.WithContext(ctx).Save(collaboration).Error; err != nil {
		return nil, err
	}
	return &collaboration, nil
}

func (r *Repo) GetCollaboration(ctx context.Context, studentId string) (*[]models.Collaboration, error) {
	var collab []models.Collaboration
	if err := r.DB.WithContext(ctx).Find(&collab, "student_id = ?", studentId).Error; err != nil {
		return nil, err
	}
	return &collab, nil
}

func (r *Repo) CreateTeam(ctx context.Context, team models.Team) (*models.Team, error) {
	if err := r.DB.WithContext(ctx).Create(team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *Repo) JoinTeam(ctx context.Context, student models.Student) (*models.Student, error) {
	if err := r.DB.WithContext(ctx).Save(student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}
