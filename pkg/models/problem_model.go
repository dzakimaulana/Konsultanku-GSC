package models

import "github.com/google/uuid"

type Problem struct {
	ID       uuid.UUID       `gorm:"primaryKey" json:"id"`
	MsmeID   string          `json:"mse_id"`
	Like     int64           `json:"like"`
	Created  int64           `json:"created"`
	Title    string          `json:"title"`
	Content  string          `json:"content"`
	Active   bool            `json:"-"`
	Tags     *[]ProblemsTags `gorm:"foreignKey:ProblemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Comments *[]Comment      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Msme     *MSME           `json:"-"`
}

type ProblemAllResp struct {
	ID           uuid.UUID      `gorm:"primaryKey" json:"id"`
	Like         int64          `json:"like"`
	CommentCount int64          `json:"comment_count"`
	Created      int64          `json:"created"`
	Title        string         `json:"title"`
	Content      string         `json:"content"`
	Msme         MSMEProbResp   `json:"msme"`
	Tags         []ProblemsTags `json:"tags"`
}

type ProblemMsmeResp struct {
	ID           uuid.UUID      `gorm:"primaryKey" json:"id"`
	Like         int64          `json:"like"`
	CommentCount int64          `json:"comment_count"`
	Created      int64          `json:"created"`
	Title        string         `json:"title"`
	Content      string         `json:"content"`
	Tags         []ProblemsTags `json:"tags"`
}
