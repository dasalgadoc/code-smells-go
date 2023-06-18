package application

import (
	"dasalgadoc.com/code_smell_go/bloaters/02-lonng_method/domain"
	shared "dasalgadoc.com/code_smell_go/bloaters/02-lonng_method/shared/domain"
)

type CourseCreator struct {
	repository domain.CourseRepository
	eventBus   shared.EventBus
}

func NewCourseCreator(repository domain.CourseRepository) *CourseCreator {
	return &CourseCreator{repository: repository}
}

func (cc CourseCreator) Invoke() error {
	course, err := domain.CreateCourse()
	if err != nil {
		return err
	}

	err = cc.repository.SaveCourse(*course)
	if err != nil {
		return err
	}

	return cc.eventBus.PublishBulk(course.PullEvents())
}
