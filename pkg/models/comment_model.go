package models

import "github.com/google/uuid"

type Comment struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	StudentID string    `json:"student_id"`
	ProblemID uuid.UUID `json:"problem_id"`
	TeamID    uuid.UUID `json:"team_id"`
	Content   string    `json:"content"`
	Team      *Team     `json:"-"`
	Problem   *Problem  `json:"-"`
	Student   *Student  `json:"-"`
}

type CommentResp struct {
	ID      uuid.UUID            `json:"id"`
	Content string               `json:"content"`
	Team    TeamResp             `json:"team"`
	Student StudentShortResponse `json:"student"`
}
