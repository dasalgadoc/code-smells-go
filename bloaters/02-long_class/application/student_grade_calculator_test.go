package application

import (
	"dasalgadoc.com/code_smell_go/bloaters/01-long_method/domain"
	"errors"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

const MAX_GRADE = 10

type studentGradeCalculatorTest struct {
	test *testing.T

	expected domain.Grades
	result   domain.Grades
	teachers domain.TeacherExtraPoint

	err         error
	expectedErr error

	studentGrade []domain.StudentGrade
}

func TestShouldReturnWeightedAverageNoExtraPoint(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	lengths := []struct {
		length int
		weight []int
	}{
		{3, []int{30, 30, 40}},
		{5, []int{10, 20, 30, 20, 20}},
		{10, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}},
	}

	for _, l := range lengths {
		s.givenASliceOfStudentGrades(l.length, l.weight)
		s.thereIsATeacherExtraPointMap()
		s.whenCalculateGrades(true, "teacher2")

		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

func TestShouldReturnWeightedAverageExtraPoint(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	lengths := []struct {
		length int
		weight []int
	}{
		{3, []int{30, 30, 40}},
		{5, []int{10, 20, 30, 20, 20}},
		{10, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}},
	}

	for _, l := range lengths {
		s.givenASliceOfStudentGrades(l.length, l.weight)
		s.thereIsATeacherExtraPointMap()
		s.whenCalculateGrades(true, "teacher1")

		s.expected += 1
		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

func TestShouldReturnZeroIfMinimumClassHasNotReached(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	lengths := []struct {
		length int
		weight []int
	}{
		{3, []int{30, 30, 40}},
		{5, []int{10, 20, 30, 20, 20}},
		{10, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}},
	}

	for _, l := range lengths {
		s.givenASliceOfStudentGrades(l.length, l.weight)
		s.thereIsATeacherExtraPointMap()
		s.whenCalculateGrades(false, "teacher1")

		s.expected = 0
		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

func TestShouldReturnZeroAtEmptyGrades(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	s.studentGrade = []domain.StudentGrade{}
	s.expected = 0

	s.whenCalculateGrades(true, "teacher1")
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")

	s.whenCalculateGrades(false, "teacher2")
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")
}

func TestShouldReturnErrorInOverWeightedGrades(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	lengths := []struct {
		length int
		weight []int
	}{
		{3, []int{40, 40, 40}},
		{5, []int{11, 20, 30, 20, 20}},
		{10, []int{20, 10, 10, 10, 10, 10, 10, 10, 10, 10}},
	}

	for _, l := range lengths {
		s.givenASliceOfStudentGrades(l.length, l.weight)
		s.andExpectedError(errors.New("the average is over-weighed"))
		s.whenCalculateGrades(true, "teacher1")

		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
		assert.Equal(t, s.expectedErr, s.err, "should return error on over-weighted grades")
		assert.Error(t, s.err, "should return error on over-weighted grades")
	}
}

func TestShouldReturnErrorInUnderWeightedGrades(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	lengths := []struct {
		length int
		weight []int
	}{
		{3, []int{20, 20, 30}},
		{5, []int{9, 20, 30, 20, 20}},
		{10, []int{0, 10, 10, 10, 10, 10, 10, 10, 10, 10}},
	}

	for _, l := range lengths {
		s.givenASliceOfStudentGrades(l.length, l.weight)
		s.andExpectedError(errors.New("the average is under-weighed"))
		s.whenCalculateGrades(true, "teacher2")

		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
		assert.Equal(t, s.expectedErr, s.err, "should return error on under-weighted grades")
		assert.Error(t, s.err, "should return error on under-weighted grades")
	}
}

/*-- steps ---*/
func startStudentGradeCalculatorTest(t *testing.T) *studentGradeCalculatorTest {
	t.Parallel()

	return &studentGradeCalculatorTest{
		test: t,
	}
}

func (s *studentGradeCalculatorTest) givenASliceOfStudentGrades(length int, weights []int) {
	var studentGrades []domain.StudentGrade
	var sumOfValues domain.Grades
	var sumOfWeights int

	for i := 0; i < length; i++ {
		number := domain.Grades(rand.Intn(MAX_GRADE + 1))
		studentGrades = append(studentGrades,
			domain.StudentGrade{
				Value:  number,
				Weight: weights[i],
			})
		sumOfValues += number * domain.Grades(weights[i]) / 100
		sumOfWeights += weights[i]
	}

	s.studentGrade = studentGrades
	s.expected = sumOfValues / domain.Grades(length)
}

func (s *studentGradeCalculatorTest) thereIsATeacherExtraPointMap() {
	s.teachers = make(domain.TeacherExtraPoint)
	s.teachers["teacher1"] = make(map[int]bool)
	s.teachers["teacher2"] = make(map[int]bool)
	startYear := time.Now().Year() - 2
	for i := startYear; i < (startYear + 5); i++ {
		s.teachers["teacher1"][i] = true
		s.teachers["teacher2"][i] = false
	}
}

func (s *studentGradeCalculatorTest) andExpectedError(err error) {
	s.expected = 0
	s.expectedErr = err
}

func (s *studentGradeCalculatorTest) whenCalculateGrades(hasReachMinimumGrades bool, teacher string) {
	calculator := NewStudentGradeCalculator(s.teachers)
	s.result, s.err = calculator.calculateGrades(s.studentGrade, hasReachMinimumGrades, teacher)
}
