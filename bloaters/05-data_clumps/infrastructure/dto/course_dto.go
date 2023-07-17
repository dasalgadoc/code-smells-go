package dto

type CourseDTO struct {
	ID       string `json:"course_id"`
	Name     string `json:"course_name"`
	Duration int    `json:"course_duration"`
}
