package msme

import (
	"context"
	"konsultanku-v2/pkg/models"
)

type AddReq struct {
	Name  string `json:"name"`
	Since string `json:"since"`
	Type  string `json:"type"`
}

type UpdateProgress struct {
	Progress    int8   `json:"progress"`
	Description string `json:"description"`
}

type EndCollaboration struct {
	Description string `json:"description"`
	Feedback    string `json:"feedback"`
	Rating      int8   `json:"rating"`
}

type MsmeRepo interface {
	AddProfile(ctx context.Context, msme models.MSME) (*models.MSME, error)
	GetByID(ctx context.Context, id string) (*models.MSME, error)
	AddedCollab(ctx context.Context, cl models.Collaboration) (*models.Collaboration, error)
	GiveProgress(ctx context.Context, collab models.Collaboration) (*models.Collaboration, error)
	StudentRating(ctx context.Context, student models.Student) (*models.Student, error)
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	FindProblemTags(ctx context.Context, id string) (*[]models.ProblemsTags, error)
}

type MsmeSvc interface {
	GetOwnProfile(ctx context.Context, id string) (*models.MSMEOwnResp, error)
	AddProfile(ctx context.Context, req AddReq, id string) error
	AddedCollab(ctx context.Context, studentId string, msmeId string) error
	GiveProgress(ctx context.Context, req UpdateProgress, msmeId string, studentId string) error
	EndCollaboration(ctx context.Context, req EndCollaboration, msmeId string, studentId string) error
}
