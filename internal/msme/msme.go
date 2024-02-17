package msme

import (
	"context"
	"konsultanku-v2/pkg/models"
)

type AddReq struct {
	Name  string `json:"name"`
	Since string `json:"since"`
}

type MsmeRepo interface {
	AddProfile(ctx context.Context, msme models.MSME) (*models.MSME, error)
	GetByID(ctx context.Context, id string) (*models.MSME, error)
}

type MsmeSvc interface {
}
