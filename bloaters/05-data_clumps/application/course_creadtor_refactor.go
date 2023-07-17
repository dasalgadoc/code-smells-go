package application

import (
	"dasalgadoc.com/code_smell_go/bloaters/05-data_clumps/domain"
	"dasalgadoc.com/code_smell_go/bloaters/05-data_clumps/infrastructure/dto"
	"github.com/gin-gonic/gin"
)

type CourseCreatorRefactor struct {
	repository domain.CourseRepository
}

func NewCourseCreatorRefactor(repository domain.CourseRepository) *CourseCreatorRefactor {
	return &CourseCreatorRefactor{repository: repository}
}

// Parallel change
func (c *CourseCreatorRefactor) Invoke(context *gin.Context) error {
	var requestBody map[string]any
	if err := context.BindJSON(&requestBody); err != nil {
		return err
	}

	return c.invokeWithParameters(requestBody["name"].(string), requestBody["duration"].(int))
}

func (c *CourseCreatorRefactor) invokeWithParameters(name string, duration int) error {
	course, err := domain.NewCourse(name, duration)
	if err != nil {
		return err
	}

	return c.repository.Save(*course)
}

func (c *CourseCreatorRefactor) InvokeStepTwo(name string, duration int) error {
	course, err := domain.NewCourse(name, duration)
	if err != nil {
		return err
	}

	return c.repository.Save(*course)
}

func (c *CourseCreatorRefactor) InvokeStepThree(courseDto dto.CourseDTO) error {
	course, err := domain.NewCourse(courseDto.Name, courseDto.Duration)
	if err != nil {
		return err
	}

	return c.repository.Save(*course)
}
