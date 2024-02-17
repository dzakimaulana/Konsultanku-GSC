package models

import "github.com/google/uuid"

type Team struct {
	ID       uuid.UUID  `gorm:"primaryKey" json:"-"`
	Name     string     `json:"name"`
	Rating   float64    `json:"rating"`
	Created  int64      `json:"-"`
	Updated  int64      `json:"-"`
	Comment  *[]Comment `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Students *[]Student `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}
