package application

import (
	"dasalgadoc.com/code_smell_go/bloaters/01-long_method/domain"
	"errors"
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

func (s *studentGradeCalculator) calculateGrades(
	examsGrades []domain.StudentGrade,
	hasReachMinimumGrades bool,
	teacher string) (domain.Grades, error) {

	var average domain.Grades
	gradesSum, gradesCount, err := s.sumGrades(examsGrades)
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

func (s *studentGradeCalculator) sumGrades(examsGrades []domain.StudentGrade) (domain.Grades, domain.Grades, error) {
	var gradesSum domain.Grades
	var gradesCount domain.Grades
	weightSum := 0
	for _, grade := range examsGrades {
		gradesSum += grade.Value * domain.Grades(grade.Weight) / 100
		weightSum += grade.Weight
		gradesCount++
	}

	err := s.checkErrorOnWeights(weightSum)
	return gradesSum, gradesCount, err
}

func (s *studentGradeCalculator) checkErrorOnWeights(weightSum int) error {
	var err error
	if weightSum > 100 {
		err = errors.New("the average is over-weighed")
	}
	if weightSum < 100 {
		err = errors.New("the average is under-weighed")
	}
	return err
}

func (s *studentGradeCalculator) hasToIncreaseOneExtraPoint(teacher string) bool {
	currentYear := time.Now().Year()
	extra, ok := s.teacherExtraPoint[teacher][currentYear]
	if ok {
		return extra
	}
	return ok
}
