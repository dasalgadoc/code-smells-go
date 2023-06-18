package domain

import (
	"dasalgadoc.com/code_smell_go/bloaters/02-/shared/domain"
	"time"
)

const EVENT_NAME = "dasalgadoc.com.code_smell.1.course.created"

type CourseCreated struct {
	fields   *domain.DomainEventCommonFields
	CourseId string
}

func (c *CourseCreated) GetDomainEventName() string {
	return c.fields.GetDomainEventName()
}

func (c *CourseCreated) GetEventId() string {
	return c.fields.GetEventId()
}

func (c *CourseCreated) GerOccurredOn() time.Time {
	return c.fields.GerOccurredOn()
}

func NewCourseCreated(courseId string) *CourseCreated {
	courseCreated := &CourseCreated{
		fields:   domain.NewDomainEventCommonFields(EVENT_NAME),
		CourseId: courseId,
	}
	return courseCreated
}
