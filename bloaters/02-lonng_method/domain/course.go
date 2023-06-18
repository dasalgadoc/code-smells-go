package domain

import "dasalgadoc.com/code_smell_go/bloaters/02-lonng_method/shared/domain"

type Course struct {
	aggregate domain.AggregateRoot
	id        *CourseId
}

func (c *Course) Id() string {
	return c.id.Value()
}

func (c *Course) PullEvents() []domain.DomainEvent {
	return c.aggregate.PullEvents()
}

func CreateCourse() (*Course, error) {
	id, err := NewCourseId()
	if err != nil {
		return nil, err
	}
	course := &Course{
		aggregate: *domain.NewAggregateRoot(),
		id:        id,
	}
	course.aggregate.RecordEvent(NewCourseCreated(course.Id()))

	return course, nil
}

func NewStudentWithId(id string) (*Course, error) {
	uid, err := NewCourseIdFromString(id)
	if err != nil {
		return nil, err
	}
	return &Course{id: uid}, nil
}
