package models

type MSME struct {
	ID            string           `gorm:"primaryKey" json:"id"`
	Name          string           `json:"name"`
	Since         string           `json:"since"`
	Type          string           `json:"type"`
	Created       int64            `json:"-"`
	Updated       int64            `json:"-"`
	Problem       *Problem         `gorm:"foreignKey:MsmeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Tags          *[]UsersTags     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Collaboration *[]Collaboration `gorm:"foreignKey:MseID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}

type MSMEOwnResp struct {
	User          UserResponse     `json:"user"`
	Name          string           `json:"name"`
	Since         string           `json:"since"`
	Problem       Problem          `json:"problem"`
	Tags          []UsersTags      `json:"tags"`
	Collaboration []CollabMsmeResp `json:"collaboration"`
}

type MSMEProbResp struct {
	User  UserResponse `json:"user"`
	Name  string       `json:"name"`
	Since string       `json:"since"`
}
