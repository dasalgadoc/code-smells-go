package domain

import (
	"dasalgadoc.com/code_smell_go/refactors/api_example/shared/domain"
)

type CourseId struct {
	value domain.UUIDValueObject
}

func (c *CourseId) Value() string {
	return c.value.Value()
}

func NewCourseId() (*CourseId, error) {
	id, err := domain.NewUUIDValueObject()
	if err != nil {
		return nil, err
	}
	return &CourseId{value: *id}, nil
}

func NewCourseIdFromString(id string) (*CourseId, error) {
	uid, err := domain.NewUUIDValueObjectFromString(id)
	if err != nil {
		return nil, err
	}
	return &CourseId{value: *uid}, nil
}
