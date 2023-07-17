package domain

import (
	shared "dasalgadoc.com/code_smell_go/refactors/api_example/shared/domain"
)

type Course struct {
	courseId shared.UUIDValueObject
	name     CourseName
	duration CourseDuration
}

func NewCourse(name string, duration int) (*Course, error) {
	id, err := shared.NewUUIDValueObject()
	if err != nil {
		return nil, err
	}
	courseName, err := NewCourseName(name)
	if err != nil {
		return nil, err
	}

	courseDuration, err := NewCourseDuration(duration)
	if err != nil {
		return nil, err
	}

	return &Course{
		courseId: *id,
		name:     *courseName,
		duration: *courseDuration,
	}, nil
}

func (c *Course) Id() string {
	return c.courseId.Value()
}

func (c *Course) Name() string {
	return c.name.Value()
}

func (c *Course) Duration() int {
	return c.duration.Value()
}
