# 👯‍♀️ Data Clumps

## 💠 This Code

## The request object

This code is a good example of a data clump. 

We decided use Gin as a web framework, and we are using the request object to create the course.
So, we send the Gin Context to inner layers, and we are coupling the infrastructure with the application and domain layers.

```go
type CourseController struct {
    courseCreator application.CourseCreator
}

func (cc *CourseController) Create(ginCtx *gin.Context) {
    err := cc.courseCreator.Invoke(ginCtx)
    // ...
}
```

__Application coupling__
```go
package application

import (
    "dasalgadoc.com/code_smell_go/bloaters/05-data-clumps/domain"
    // A gin dependency in application layer its a red flag
    "github.com/gin-gonic/gin"
)

type CourseCreator struct {
    repository domain.CourseRepository
}

// Really bad and coupled with gin
func (c *CourseCreator) Invoke(context *gin.Context) error {
    var requestBody map[string]any
	// Figure we don't know how to parse the request body, so we have a generic solution
    if err := context.BindJSON(&requestBody); err != nil {
        return err
    }
	
    course, err := domain.NewCourse(requestBody["name"].(string), requestBody["duration"].(int))
    if err != nil {
        return err
    }
	
    return c.repository.Save(*course)
}

```

Here we have a data clump, passing a data through methods, class and structs; generating coupling, and if Gin changes (or We want to changes) we have to modify inner layers.

## 🧑🏻‍🔬 Refactoring

__Step 1: Separated the application from concrete infrastructure__

```go
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
```

__Step 2: Bring infrastructure code to infrastructure__

```go
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
```

```go
func (c *CourseCreatorRefactor) InvokeStepTwo(name string, duration int) error {
    course, err := domain.NewCourse(name, duration)
    if err != nil {
        return err
    }

    return c.repository.Save(*course)
}
```
But, what happens if a request has more parameters? We can return to Long parameter list code smell, so we can use DTO structs.

__Step 3: Use DTOs__

```go
type CourseDTO struct {
    ID       string `json:"course_id"`
    Name     string `json:"course_name"`
    Duration int    `json:"course_duration"`
}
```

```go
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
```

```go
func (c *CourseCreatorRefactor) InvokeStepThree(courseDto dto.CourseDTO) error {
    course, err := domain.NewCourse(courseDto.Name, courseDto.Duration)
    if err != nil {
        return err
    }

    return c.repository.Save(*course)
}
```