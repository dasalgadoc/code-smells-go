package application

import (
	"dasalgadoc.com/code_smell_go/bloaters/01-long_method/domain"
	"time"
)

type studentGradeCalculatorRefactor struct {
	teacherExtraPoint domain.TeacherExtraPoint
}

func (s *studentGradeCalculatorRefactor) calculateGradesMinimumClassesAndWeightedAverageWithExtraPoint(
	examsGrades []domain.StudentGrade,
	hasReachMinimumGrades bool,
	teacher string) domain.Grades {
	if !(len(examsGrades) == 0) {
		var gradesSum domain.Grades
		var gradesCount domain.Grades
		weightSum := 0

		for _, grade := range examsGrades {
			gradesSum += grade.Value * domain.Grades(grade.Weight) / 100
			weightSum += grade.Weight
			gradesCount++
		}

		// errors handling: oue return type is grades, so we can't return an error
		// magic numbers'
		// concept overload
		if weightSum == 100 {
			if hasReachMinimumGrades {
				if s.hasToIncreaseOneExtraPoint(teacher) {
					return gradesSum/gradesCount + 1
				}
				return gradesSum / gradesCount
			} else {
				return 0
			}
		} else if weightSum > 100 {
			return -1 // error type over-weighed
		} else {
			return -2 // error type under-weighed
		}

	} else {
		return 0
	}
}

func (s *studentGradeCalculatorRefactor) hasToIncreaseOneExtraPoint(teacher string) bool {
	var gotExtraPoint bool
	if s.teacherExtraPoint != nil {
		currentYear := time.Now().Year()
		if extra, ok := s.teacherExtraPoint[teacher][currentYear]; ok {
			gotExtraPoint = extra
		}
	}
	return gotExtraPoint
}
