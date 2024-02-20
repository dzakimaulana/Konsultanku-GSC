package models

import "github.com/google/uuid"

type Student struct {
	ID            string           `gorm:"primaryKey" json:"id"`
	TeamID        uuid.UUID        `json:"team_id"`
	Major         string           `json:"major"`
	ClassOf       string           `json:"class_of"`
	University    string           `json:"university"`
	Rating        float32          `json:"rating"`
	Created       int64            `json:"created"`
	Updated       int64            `json:"updated"`
	Collaboration *[]Collaboration `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Tags          *[]UsersTags     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Comment       *[]Comment       `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Team          *Team            `json:"-"`
}

type StudentResponse struct {
	User          UserResponse        `json:"user"`
	Major         string              `json:"major"`
	ClassOf       string              `json:"class_of"`
	University    string              `json:"university"`
	Tags          []UsersTags         `json:"tags"`
	Team          Team                `json:"team"`
	Collaboration []CollabStudentResp `json:"collaboration"`
}

type StudentShortResponse struct {
	User       UserResponse `json:"user"`
	Major      string       `json:"major"`
	ClassOf    string       `json:"class_of"`
	University string       `json:"university"`
}
