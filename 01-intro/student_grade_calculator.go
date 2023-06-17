package _1_intro

import "time"

type studentGradeCalculator struct {
	teacherExtraPoint teacherExtraPoint
}

func NewStudentGradeCalculator(point teacherExtraPoint) *studentGradeCalculator {
	return &studentGradeCalculator{
		teacherExtraPoint: point,
	}
}

func (s *studentGradeCalculator) calculateGrades(examsGrades []grades) grades {
	if !(len(examsGrades) == 0) {
		var gradesSum grades
		var gradesCount grades

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
	examsGrades []grades,
	hasReachMinimumGrades bool) grades {
	if !(len(examsGrades) == 0) {
		var gradesSum grades
		var gradesCount grades

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
	examsGrades []StudentGrade,
	hasReachMinimumGrades bool) grades {
	if !(len(examsGrades) == 0) {
		var gradesSum grades
		var gradesCount grades
		weightSum := 0

		for _, grade := range examsGrades {
			gradesSum += grade.value * grades(grade.weight) / 100
			weightSum += grade.weight
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
	examsGrades []StudentGrade,
	hasReachMinimumGrades bool,
	teacher string) grades {
	if !(len(examsGrades) == 0) {
		var gradesSum grades
		var gradesCount grades
		weightSum := 0
		gotExtraPoint := false

		if s.teacherExtraPoint != nil {
			currentYear := time.Now().Year()
			if extra, ok := s.teacherExtraPoint[teacher][currentYear]; ok {
				gotExtraPoint = extra
			}
		}

		for _, grade := range examsGrades {
			gradesSum += grade.value * grades(grade.weight) / 100
			weightSum += grade.weight
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
