package controller

import (
	"dasalgadoc.com/code_smell_go/bloaters/05-data-clumps/application"
	"dasalgadoc.com/code_smell_go/bloaters/05-data-clumps/infrastructure/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CourseControllerRefactor struct {
	courseCreator application.CourseCreatorRefactor
}

func NewCourseControllerRefactor(courseCreator application.CourseCreatorRefactor) *CourseControllerRefactor {
	return &CourseControllerRefactor{courseCreator: courseCreator}
}

func (cc *CourseControllerRefactor) CreateStepOne(ginCtx *gin.Context) {
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

func (cc *CourseControllerRefactor) CreateStepTwo(ginCtx *gin.Context) {
	var requestBody map[string]any
	if err := ginCtx.BindJSON(&requestBody); err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	err := cc.courseCreator.InvokeStepTwo(requestBody["name"].(string), requestBody["duration"].(int))
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ginCtx.JSON(http.StatusCreated, gin.H{
		"message": "Course created successfully",
	})
}

func (cc *CourseControllerRefactor) CreateStepThree(ginCtx *gin.Context) {
	var request dto.CourseDTO
	if err := ginCtx.BindJSON(&request); err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	err := cc.courseCreator.InvokeStepThree(request)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ginCtx.JSON(http.StatusCreated, gin.H{
		"message": "Course created successfully",
	})
}
