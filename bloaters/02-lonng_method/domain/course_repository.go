package domain

type CourseRepository interface {
	GetCourseById(id CourseId) (Course, error)
	SaveCourse(course Course) error
}
