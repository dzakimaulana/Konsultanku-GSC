package models

import "github.com/google/uuid"

type Comment struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	ProblemID uuid.UUID `json:"problem_id"`
	TeamID    uuid.UUID `json:"team_id"`
	Content   string    `json:"content"`
	Team      *Team     `json:"-"`
	Problem   *Problem  `json:"-"`
}
