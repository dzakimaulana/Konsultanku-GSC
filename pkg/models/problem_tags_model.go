package models

type ProblemsTags struct {
	ProblemID string `json:"-"`
	TagID     int64  `json:"-"`
	Tag       *Tags  `json:"tag"`
}
