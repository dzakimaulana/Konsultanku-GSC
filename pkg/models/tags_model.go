package models

type Tags struct {
	ID    int64        `gorm:"primaryKey" json:"id"`
	Name  string       `json:"name"`
	Users *[]UsersTags `gorm:"foreignKey:TagsID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}
