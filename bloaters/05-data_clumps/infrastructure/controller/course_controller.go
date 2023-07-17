package controller

import (
	"dasalgadoc.com/code_smell_go/bloaters/05-data_clumps/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CourseController struct {
	courseCreator application.CourseCreator
}

func NewCourseController(courseCreator application.CourseCreator) *CourseController {
	return &CourseController{courseCreator: courseCreator}
}

func (cc *CourseController) Create(ginCtx *gin.Context) {
	err := cc.courseCreator.Invoke(ginCtx)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ginCtx.JSON(http.StatusCreated, gin.H{
		"message": "Course created successfully",
	})
}
