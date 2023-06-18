package application

import (
	"dasalgadoc.com/code_smell_go/bloaters/01-long_method/domain"
	"time"
)

type studentGradeCalculator struct {
	teacherExtraPoint domain.TeacherExtraPoint
}

func NewStudentGradeCalculator(point domain.TeacherExtraPoint) *studentGradeCalculator {
	return &studentGradeCalculator{
		teacherExtraPoint: point,
	}
}

func (s *studentGradeCalculator) calculateGrades(examsGrades []domain.Grades) domain.Grades {
	if !(len(examsGrades) == 0) {
		var gradesSum domain.Grades
		var gradesCount domain.Grades

		for _, grade := range examsGrades {
			gradesSum += grade
			gradesCount++
		}

		return gradesSum / gradesCount
	} else {
		return 0
	}
}

func (s *studentGradeCalculator) calculateGradesMinimumClasses(
	examsGrades []domain.Grades,
	hasReachMinimumGrades bool) domain.Grades {
	if !(len(examsGrades) == 0) {
		var gradesSum domain.Grades
		var gradesCount domain.Grades

		for _, grade := range examsGrades {
			gradesSum += grade
			gradesCount++
		}

		if hasReachMinimumGrades {
			return gradesSum / gradesCount
		} else {
			return 0
		}

	} else {
		return 0
	}
}
func (s *studentGradeCalculator) calculateGradesMinimumClassesAndWeightedAverage(
	examsGrades []domain.StudentGrade,
	hasReachMinimumGrades bool) domain.Grades {
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

func (s *studentGradeCalculator) calculateGradesMinimumClassesAndWeightedAverageWithExtraPoint(
	examsGrades []domain.StudentGrade,
	hasReachMinimumGrades bool,
	teacher string) domain.Grades {
	if !(len(examsGrades) == 0) {
		var gradesSum domain.Grades
		var gradesCount domain.Grades
		weightSum := 0
		gotExtraPoint := false

		if s.teacherExtraPoint != nil {
			currentYear := time.Now().Year()
			if extra, ok := s.teacherExtraPoint[teacher][currentYear]; ok {
				gotExtraPoint = extra
			}
		}

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
				if gotExtraPoint {
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
