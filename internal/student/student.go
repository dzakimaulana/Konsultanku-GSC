package student

import (
	"context"
	"konsultanku-v2/pkg/models"
)

type AddStudent struct {
	Major      string `json:"major"`
	ClassOf    string `json:"class_of"`
	University string `json:"university"`
}

type CreateTeam struct {
	Name string `json:"name"`
}

type StudentRepo interface {
	AddProfile(ctx context.Context, student models.Student) (*models.Student, error)
	GetByID(ctx context.Context, id string) (*models.Student, error)
	UpdateCollaboration(ctx context.Context, collaboration models.Collaboration) (*models.Collaboration, error)
	GetCollaboration(ctx context.Context, studentId string) (*[]models.Collaboration, error)
	CreateTeam(ctx context.Context, team models.Team) (*models.Team, error)
	JoinTeam(ctx context.Context, student models.Student) (*models.Student, error)
}

type StudentSvc interface {
	AddProfile(ctx context.Context, studentReq AddStudent, id string) error
	GetOwnProfile(ctx context.Context, id string) (*models.StudentResponse, error)
	AcceptOffer(ctx context.Context, msmeId string, studentId string) (*models.UserPhoneResp, error)
	GetCollaboration(ctx context.Context, studentId string) (*[]models.GetStudentCollaboration, error)
	CreateTeam(ctx context.Context, req CreateTeam) (*models.TeamResp, error)
	JoinTeam(ctx context.Context, teamId string, studentId string) error
}
