package domain

type CourseRepository interface {
	Save(course Course) error
}
