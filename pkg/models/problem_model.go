package models

import "github.com/google/uuid"

type Problem struct {
	ID       uuid.UUID  `gorm:"primaryKey" json:"id"`
	MsmeID   string     `json:"mse_id"`
	Like     int64      `json:"like"`
	Created  int64      `json:"created"`
	Content  string     `json:"content"`
	Active   bool       `json:"-"`
	Comments *[]Comment `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Msme     *MSME      `json:"-"`
}

type ProblemAllRes struct {
	ID      uuid.UUID `gorm:"primaryKey" json:"id"`
	Like    int64     `json:"like"`
	Created int64     `json:"created"`
	Content string    `json:"content"`
	Msme    MSME      `json:"msme"`
}

type ProblemIDRes struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	Like     int64     `json:"like"`
	Created  int64     `json:"created"`
	Content  string    `json:"content"`
	Comments []Comment `json:"comments"`
	Msme     MSME      `json:"msme"`
}
