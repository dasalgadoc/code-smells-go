package dto

type Course struct {
	courseType string  `json:"course_type"`
	duration   float32 `json:"duration"`
	points     float32 `json:"points"`
}
