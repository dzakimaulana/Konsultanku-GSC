package student

import (
	"context"
	"konsultanku-v2/pkg/models"
)

type AddStudent struct {
	Name       string `json:"name"`
	Major      string `json:"major"`
	ClassOf    string `json:"class_of"`
	University string `json:"university"`
}

type StudentRepo interface {
	GetByID(ctx context.Context, id string) (*models.Student, error)
}

type StudentSvc interface {
}
