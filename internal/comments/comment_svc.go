package comments

import (
	"context"
	"konsultanku-v2/pkg/databases"
	"konsultanku-v2/pkg/models"
	"time"

	"github.com/google/uuid"
)

type Svc struct {
	CommentRepo
	timeout time.Duration
}

func NewSvc(cr CommentRepo) *Svc {
	return &Svc{
		CommentRepo: cr,
		timeout:     time.Duration(10) * time.Second,
	}
}

func (s *Svc) AddComment(ctx context.Context, commentReq AddComment, studentId string, teamId string, problemId string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	probId, err := uuid.Parse(problemId)
	if err != nil {
		return err
	}
	var comment *models.Comment
	if teamId == "" {
		comment = &models.Comment{
			ID:        uuid.New(),
			StudentID: studentId,
			ProblemID: probId,
			Content:   commentReq.Content,
		}
	} else {
		ti, err := uuid.Parse(teamId)
		if err != nil {
			return err
		}
		comment = &models.Comment{
			ID:        uuid.New(),
			StudentID: studentId,
			ProblemID: probId,
			TeamID:    ti,
			Content:   commentReq.Content,
		}
	}

	_, err = s.CommentRepo.AddComment(ctx, *comment)
	if err != nil {
		return err
	}
	return nil
}

func (s *Svc) GetComments(ctx context.Context, mseId string) (*[]models.CommentResp, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	resp, err := s.CommentRepo.GetComments(ctx, mseId)
	if err != nil {
		return nil, err
	}
	var commentsResp []models.CommentResp
	for _, comment := range *resp {
		var teamResp *models.TeamResp
		if comment.Team == nil {
			teamResp = &models.TeamResp{
				ID:     comment.Team.ID.String(),
				Name:   comment.Team.Name,
				Rating: float64(comment.Team.Rating),
			}
		}
		studentId := comment.Student.ID
		userResp, err := databases.AuthMd.GetUser(ctx, studentId)
		if err != nil {
			return nil, err
		}
		user := &models.UserResponse{
			UID:         userResp.UID,
			Email:       userResp.Email,
			DisplayName: userResp.DisplayName,
			PhoneNumber: userResp.PhoneNumber,
			PhotoURL:    userResp.PhotoURL,
		}

		student := &models.StudentShortResponse{
			User:       *user,
			Major:      comment.Student.Major,
			University: comment.Student.University,
			ClassOf:    comment.Student.ClassOf,
		}

		commentsResp = append(commentsResp, models.CommentResp{
			ID:      comment.ID,
			Content: comment.Content,
			Team:    *teamResp,
			Student: *student,
		})
	}
	return &commentsResp, nil
}
