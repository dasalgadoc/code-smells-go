package application

import (
	"dasalgadoc.com/code_smell_go/bloaters/01-long_method/domain"
	"errors"
	"time"
)

type studentGradeCalculatorRefactor struct {
	teacherExtraPoint domain.TeacherExtraPoint
}

func NewStudentGradeCalculatorRefactor(point domain.TeacherExtraPoint) *studentGradeCalculatorRefactor {
	return &studentGradeCalculatorRefactor{
		teacherExtraPoint: point,
	}
}

func (s *studentGradeCalculatorRefactor) calculateGradesRefactorOne(
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

func (s *studentGradeCalculatorRefactor) calculateGradesRefactorTwo(
	examsGrades []domain.StudentGrade,
	hasReachMinimumGrades bool,
	teacher string) domain.Grades {
	if !(len(examsGrades) == 0) {
		gradesSum, gradesCount, weightSum := s.sumGradesAndWeights(examsGrades)

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

// Guard clauses
func (s *studentGradeCalculatorRefactor) calculateGradesRefactorThree(
	examsGrades []domain.StudentGrade,
	hasReachMinimumGrades bool,
	teacher string) domain.Grades {
	if len(examsGrades) == 0 {
		return 0
	}

	gradesSum, gradesCount, weightSum := s.sumGradesAndWeights(examsGrades)

	// errors handling: oue return type is grades, so we can't return an error
	// magic numbers'
	// concept overload
	if weightSum > 100 {
		return -1
	}
	if weightSum < 100 {
		return -2
	}
	if !hasReachMinimumGrades {
		return 0
	}
	if s.hasToIncreaseOneExtraPoint(teacher) {
		return gradesSum/gradesCount + 1
	}

	return gradesSum / gradesCount
}

// Specify error and return types
func (s *studentGradeCalculatorRefactor) calculateGradesRefactorFour(
	examsGrades []domain.StudentGrade,
	hasReachMinimumGrades bool,
	teacher string) (domain.Grades, error) {

	var average domain.Grades
	if len(examsGrades) == 0 {
		return average, errors.New("there is no grades")
	}

	gradesSum, gradesCount, err := s.sumGradesAndWeightsWithErrors(examsGrades)
	if err != nil {
		return average, err
	}
	if !hasReachMinimumGrades {
		return average, nil
	}

	average = gradesSum / gradesCount
	if s.hasToIncreaseOneExtraPoint(teacher) {
		average += 1
	}

	return average, nil
}

// Simplify error with domain knowledge
func (s *studentGradeCalculatorRefactor) calculateGradesRefactorFive(
	examsGrades []domain.StudentGrade,
	hasReachMinimumGrades bool,
	teacher string) (domain.Grades, error) {
	var average domain.Grades
	gradesSum, gradesCount, err := s.sumGradesAndWeightsWithErrors(examsGrades)
	if err != nil {
		return average, err
	}
	if !hasReachMinimumGrades {
		return average, nil
	}

	average = gradesSum / gradesCount
	if s.hasToIncreaseOneExtraPoint(teacher) {
		average += 1
	}

	return average, nil
}

func (s *studentGradeCalculatorRefactor) sumGradesAndWeights(
	examsGrades []domain.StudentGrade) (domain.Grades, domain.Grades, int) {
	var gradesSum domain.Grades
	var gradesCount domain.Grades
	weightSum := 0

	for _, grade := range examsGrades {
		gradesSum += grade.Value * domain.Grades(grade.Weight) / 100
		weightSum += grade.Weight
		gradesCount++
	}
	return gradesSum, gradesCount, weightSum
}

func (s *studentGradeCalculatorRefactor) sumGradesAndWeightsWithErrors(
	examsGrades []domain.StudentGrade) (domain.Grades, domain.Grades, error) {
	var gradesSum domain.Grades
	var gradesCount domain.Grades
	weightSum := 0
	for _, grade := range examsGrades {
		gradesSum += grade.Value * domain.Grades(grade.Weight) / 100
		weightSum += grade.Weight
		gradesCount++
	}

	err := s.checkErrorsOnAverageCollection(weightSum)
	return gradesSum, gradesCount, err
}

func (s *studentGradeCalculatorRefactor) checkErrorsOnAverageCollection(weightSum int) error {
	var err error
	if weightSum > 100 {
		err = errors.New("the average is over-weighed")
	}
	if weightSum < 100 {
		err = errors.New("the average is over-weighed")
	}
	return err
}

func (s *studentGradeCalculatorRefactor) hasToIncreaseOneExtraPoint(teacher string) bool {
	currentYear := time.Now().Year()
	extra, ok := s.teacherExtraPoint[teacher][currentYear]
	if ok {
		return extra
	}
	return ok
}
