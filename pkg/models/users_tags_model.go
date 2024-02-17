package models

type UsersTags struct {
	UserID string `json:"-"`
	TagID  int64  `json:"-"`
	Tag    *Tags  `json:"tag"`
}
