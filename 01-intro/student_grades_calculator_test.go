package _1_intro

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

const MAX_GRADE = 10

type studentGradeCalculatorTest struct {
	test     *testing.T
	grades   []grades
	expected grades
	result   grades

	teachers teacherExtraPoint

	studentGrade []StudentGrade
}

/*-- Feature 1: average --*/
func TestShouldReturnAverageOnASlicesOfNumber(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	lengths := []int{3, 5, 10, 20, 50, 100}
	for _, l := range lengths {
		s.givenASliceOfNumbers(l)
		s.whenCalculateGrades()

		assert.Equal(t, s.expected, s.result, "should return average on a slice of numbers")
	}
}

func TestShouldReturnZeroOnEmptySlice(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	s.grades = []grades{}
	s.expected = 0

	s.whenCalculateGrades()
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")
}

/*-- Feature 2: minimum --*/
func TestShouldReturnAverageOnASliceNumberWithMinimumClassReached(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	lengths := []int{3, 5, 10, 20, 50, 100}
	for _, l := range lengths {
		s.givenASliceOfNumbers(l)
		s.whenCalculateGradesMinimumReached(true)

		assert.Equal(t, s.expected, s.result, "should return average on a slice of numbers")
	}
}

func TestShouldReturnZeroOnASliceNumberWithoutMinimumClassReached(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	lengths := []int{3, 5, 10, 20, 50, 100}
	for _, l := range lengths {
		s.givenASliceOfNumbers(l)
		s.whenCalculateGradesMinimumReached(false)

		assert.Equal(t, grades(0), s.result, "should return average on a slice of numbers")
	}
}

func TestShouldReturnZeroOnEmptySliceMinimum(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	s.grades = []grades{}
	s.expected = 0

	s.whenCalculateGradesMinimumReached(true)
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")

	s.whenCalculateGradesMinimumReached(false)
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")
}

/*-- Feature 3: weighted --*/
func TestShouldReturnWeightedAverageOnASliceNumberWithMinimumClassReached(t *testing.T) {
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
		s.whenCalculateGradesMinimumReachedAndWeightedAverage(true)

		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

func TestShouldReturnZeroOnASliceNumberWithoutMinimumClassReachedWeighted(t *testing.T) {
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
		s.whenCalculateGradesMinimumReachedAndWeightedAverage(false)

		assert.Equal(t, grades(0), s.result, "should return average on a slice of grades")
	}
}

func TestShouldReturnZeroOnEmptySliceMinimumWeighted(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	s.studentGrade = []StudentGrade{}
	s.expected = 0

	s.whenCalculateGradesMinimumReachedAndWeightedAverage(true)
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")

	s.whenCalculateGradesMinimumReachedAndWeightedAverage(false)
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")
}

func TestShouldReturnAnErrorCodeWithWWeightedUnderOneHundred(t *testing.T) {
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
		s.whenCalculateGradesMinimumReachedAndWeightedAverage(true)

		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

func TestShouldReturnAnErrorCodeWithWWeightedOverOneHundred(t *testing.T) {
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
		s.whenCalculateGradesMinimumReachedAndWeightedAverage(true)

		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

/*-- Feature 4: weighted with teachers --*/
func TestShouldReturnWeightedAverageOnASliceNumberWithMinimumClassReachedWithTeachersListNoExtraPoint(t *testing.T) {
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
		s.andThereIsAtTeacherExtraPointMap()
		s.whenCalculateGradesMinimumReachedAndWeightedAverageAndTeachers(true, "teacher2")

		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

func TestShouldReturnWeightedAverageOnASliceNumberWithMinimumClassReachedWithTeachersListExtraPoint(t *testing.T) {
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
		s.andThereIsAtTeacherExtraPointMap()
		s.whenCalculateGradesMinimumReachedAndWeightedAverageAndTeachers(true, "teacher1")

		s.expected += 1
		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

// Upper test repeated with teachers list...

/*-- steps ---*/
func startStudentGradeCalculatorTest(t *testing.T) *studentGradeCalculatorTest {
	t.Parallel()

	return &studentGradeCalculatorTest{
		test: t,
	}
}

func (s *studentGradeCalculatorTest) givenASliceOfNumbers(length int) {
	var numbers []grades
	var sumOfValues grades

	for i := 0; i < length; i++ {
		numbers = append(numbers, grades(rand.Int()))
		sumOfValues += numbers[i]
	}

	s.grades = numbers
	s.expected = sumOfValues / grades(length)
}

func (s *studentGradeCalculatorTest) givenASliceOfStudentGrades(length int, weights []int) {
	var studentGrades []StudentGrade
	var sumOfValues grades
	var sumOfWeights int

	for i := 0; i < length; i++ {
		number := grades(rand.Intn(MAX_GRADE + 1))
		studentGrades = append(studentGrades,
			StudentGrade{
				value:  number,
				weight: weights[i],
			})
		sumOfValues += number * grades(weights[i]) / 100
		sumOfWeights += weights[i]
	}

	s.studentGrade = studentGrades
	s.expected = sumOfValues / grades(length)

	if sumOfWeights > 100 {
		s.expected = -1
	}
	if sumOfWeights < 100 {
		s.expected = -2
	}
}

func (s *studentGradeCalculatorTest) andThereIsAtTeacherExtraPointMap() {
	s.teachers = make(teacherExtraPoint)
	s.teachers["teacher1"] = make(map[int]bool)
	s.teachers["teacher2"] = make(map[int]bool)
	startYear := time.Now().Year() - 2
	for i := startYear; i < (startYear + 5); i++ {
		s.teachers["teacher1"][i] = true
		s.teachers["teacher2"][i] = false
	}
}

func (s *studentGradeCalculatorTest) whenCalculateGrades() {
	target := studentGradeCalculator{}
	s.result = target.calculateGrades(s.grades)
}

func (s *studentGradeCalculatorTest) whenCalculateGradesMinimumReached(hasReachMinimum bool) {
	target := studentGradeCalculator{}
	s.result = target.calculateGradesMinimumClasses(s.grades, hasReachMinimum)
}

func (s *studentGradeCalculatorTest) whenCalculateGradesMinimumReachedAndWeightedAverage(hasReachMinimum bool) {
	target := studentGradeCalculator{}
	s.result = target.calculateGradesMinimumClassesAndWeightedAverage(s.studentGrade, hasReachMinimum)
}

func (s *studentGradeCalculatorTest) whenCalculateGradesMinimumReachedAndWeightedAverageAndTeachers(
	hasReachMinimum bool, teacher string) {
	target := NewStudentGradeCalculator(s.teachers)
	s.result = target.calculateGradesMinimumClassesAndWeightedAverageWithExtraPoint(s.studentGrade, hasReachMinimum, teacher)
}
