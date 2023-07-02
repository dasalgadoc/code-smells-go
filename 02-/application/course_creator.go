package application

import (
	domain2 "dasalgadoc.com/code_smell_go/02-/domain"
	shared "dasalgadoc.com/code_smell_go/02-/shared/domain"
)

type CourseCreator struct {
	repository domain2.CourseRepository
	eventBus   shared.EventBus
}

func NewCourseCreator(repository domain2.CourseRepository) *CourseCreator {
	return &CourseCreator{repository: repository}
}

func (cc CourseCreator) Invoke() error {
	course, err := domain2.CreateCourse()
	if err != nil {
		return err
	}

	err = cc.repository.SaveCourse(*course)
	if err != nil {
		return err
	}

	return cc.eventBus.PublishBulk(course.PullEvents())
}
