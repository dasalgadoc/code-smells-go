package infrastructure

import (
	"dasalgadoc.com/code_smell_go/change_preventers/intro/domain"
	"dasalgadoc.com/code_smell_go/change_preventers/intro/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

type stepsCalculatorControllerTest struct {
	test *testing.T

	importerMMock *test.ImporterMock

	result       *string
	importResult *domain.Table
	err          error
	expectedErr  error
}

func TestShouldReturnEmptyStepList(t *testing.T) {
	s := NewStepsCalculatorControllerTest(t)

	s.importResult = nil
	s.expectedErr = nil

	courseId := "course-1"
	s.givenAReturnFromImporter(courseId)
	s.whenInvokeStepsCalculatorController(courseId)

	assert.Equal(s.test, "[]", *s.result)
	assert.Equal(s.test, s.expectedErr, s.err)

}

func TestShouldReturnExistingCourseSteps(t *testing.T) {
	s := NewStepsCalculatorControllerTest(t)

	s.importResult = &domain.Table{
		Rows: [][]string{
			{"1", "video", "0", "13"},
			{"2", "quiz", "5", "0"},
		},
	}
	s.expectedErr = nil

	courseId := "course-1"
	s.givenAReturnFromImporter(courseId)
	s.whenInvokeStepsCalculatorController(courseId)
	assert.Equal(s.test, `[{"id":"1","type":"video","duration":14.3,"points":1430},{"id":"2","type":"quiz","duration":2.5,"points":25}]`, *s.result)
}

func TestShouldIgnoreStepWithInvalidType(t *testing.T) {
	s := NewStepsCalculatorControllerTest(t)

	s.importResult = &domain.Table{
		Rows: [][]string{
			{"1", "survey", "0", "0"},
		},
	}
	s.expectedErr = nil

	courseId := "course-1"
	s.givenAReturnFromImporter(courseId)
	s.whenInvokeStepsCalculatorController(courseId)

	assert.Equal(s.test, "[]", *s.result)
}

/*-- steps --*/
func NewStepsCalculatorControllerTest(t *testing.T) *stepsCalculatorControllerTest {
	t.Parallel()

	return &stepsCalculatorControllerTest{
		importerMMock: new(test.ImporterMock),
		test:          t,
	}
}

func (s *stepsCalculatorControllerTest) givenAReturnFromImporter(argument string) {
	s.importerMMock.On("Invoke", argument).
		Return(s.importResult, s.expectedErr).Once()
}

func (s *stepsCalculatorControllerTest) whenInvokeStepsCalculatorController(argument string) {
	target := NewStepsCalculatorController(s.importerMMock)
	s.result, s.err = target.Get(argument)
}
