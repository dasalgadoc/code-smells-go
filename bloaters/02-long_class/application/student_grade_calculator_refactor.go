package application

import (
	"dasalgadoc.com/code_smell_go/bloaters/02-long_class/domain"
	"time"
)

var CurrentYear = func() int {
	return time.Now().Year()
}

type studentGradeCalculatorRefactor struct {
	teachersExtraPoint domain.TeachersExtraPoint
}

func NewStudentGradeCalculatorRefactor(point domain.TeachersExtraPoint) *studentGradeCalculatorRefactor {
	return &studentGradeCalculatorRefactor{
		teachersExtraPoint: point,
	}
}

func (s *studentGradeCalculatorRefactor) Do(
	grades []domain.NumericGrade,
	weights []int,
	hasReachMinimumGrades bool,
	teacher string) (domain.NumericGrade, error) {
	if !hasReachMinimumGrades {
		return 0, nil
	}

	g, err := domain.NewGrades(grades, weights)
	if err != nil {
		return 0, err
	}

	return g.CalculateAverage() + s.teachersExtraPoint.GetExtraPoint(teacher, CurrentYear()), nil
}
