package domain

type CourseRepository interface {
	GetCourseById(id CourseId) (Course, error)
	CreateCourse(course Course) error
}
