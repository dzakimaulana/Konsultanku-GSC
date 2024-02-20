package comments

import (
	"context"
	"konsultanku-v2/pkg/models"
)

type AddComment struct {
	ProblemID string `json:"problem_id"`
	Content   string `json:"content"`
}

type CommentRepo interface {
	AddComment(ctx context.Context, comment models.Comment) (*models.Comment, error)
	GetComments(ctx context.Context, mseId string) (*[]models.Comment, error)
}

type CommentSvc interface {
	AddComment(ctx context.Context, commentReq AddComment, studentId string, teamId string) error
	GetComments(ctx context.Context, mseId string) (*[]models.CommentResp, error)
}
