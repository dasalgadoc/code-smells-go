package dto

type CourseJsonDto struct {
	Id         string  `json:"id"`
	CourseType string  `json:"type"`
	Duration   float64 `json:"duration"`
	Points     float64 `json:"points"`
}
