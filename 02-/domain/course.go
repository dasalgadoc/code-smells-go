package domain

import (
	domain2 "dasalgadoc.com/code_smell_go/02-/shared/domain"
)

type Course struct {
	aggregate domain2.AggregateRoot
	id        *CourseId
}

func (c *Course) Id() string {
	return c.id.Value()
}

func (c *Course) PullEvents() []domain2.DomainEvent {
	return c.aggregate.PullEvents()
}

func CreateCourse() (*Course, error) {
	id, err := NewCourseId()
	if err != nil {
		return nil, err
	}
	course := &Course{
		aggregate: *domain2.NewAggregateRoot(),
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
