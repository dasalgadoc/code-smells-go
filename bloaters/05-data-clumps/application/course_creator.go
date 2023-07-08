package application

import (
	"dasalgadoc.com/code_smell_go/bloaters/05-data-clumps/domain"
	"github.com/gin-gonic/gin"
)

type CourseCreator struct {
	repository domain.CourseRepository
}

func NewCourseCreator(repository domain.CourseRepository) *CourseCreator {
	return &CourseCreator{repository: repository}
}

// Really bad and coupled with gin
func (c *CourseCreator) Invoke(context *gin.Context) error {
	var requestBody map[string]any
	if err := context.BindJSON(&requestBody); err != nil {
		return err
	}

	course, err := domain.NewCourse(requestBody["name"].(string), requestBody["duration"].(int))
	if err != nil {
		return err
	}

	return c.repository.Save(*course)
}
