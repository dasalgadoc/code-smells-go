package application

import (
	"dasalgadoc.com/code_smell_go/bloaters/01-long_method/domain"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

const MAX_GRADE = 10

type studentGradeCalculatorTest struct {
	test     *testing.T
	grades   []domain.Grades
	expected domain.Grades
	result   domain.Grades

	teachers domain.TeacherExtraPoint

	studentGrade []domain.StudentGrade
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

	s.grades = []domain.Grades{}
	s.expected = 0

	s.whenCalculateGrades()
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")
}

/*-- Feature 2: minimum --*/
func TestShouldReturnAverageOnASliceNumberIfMinimumClassHasBeenReached(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	lengths := []int{3, 5, 10, 20, 50, 100}
	for _, l := range lengths {
		s.givenASliceOfNumbers(l)
		s.whenCalculateGradesMinimumReached(true)

		assert.Equal(t, s.expected, s.result, "should return average on a slice of numbers")
	}
}

func TestShouldReturnZeroOnASliceOfNumbersIfMinimumClassHasNoBeenReached(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	lengths := []int{3, 5, 10, 20, 50, 100}
	for _, l := range lengths {
		s.givenASliceOfNumbers(l)
		s.whenCalculateGradesMinimumReached(false)

		assert.Equal(t, domain.Grades(0), s.result, "should return average on a slice of numbers")
	}
}

func TestShouldReturnZeroOnEmptySliceWithAndWithoutMinimumReached(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	s.grades = []domain.Grades{}
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

		assert.Equal(t, domain.Grades(0), s.result, "should return average on a slice of grades")
	}
}

func TestShouldReturnZeroOnEmptySliceWeightedStructure(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	s.studentGrade = []domain.StudentGrade{}
	s.expected = 0

	s.whenCalculateGradesMinimumReachedAndWeightedAverage(true)
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")

	s.whenCalculateGradesMinimumReachedAndWeightedAverage(false)
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")
}

func TestShouldReturnAnErrorCodeWithWeightedUnderOneHundred(t *testing.T) {
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

func TestShouldReturnAnErrorCodeWithWeightedOverOneHundred(t *testing.T) {
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
func TestShouldReturnWeightedAverageOnASliceNumberWithMinimumClassReachedWithTeachersListNoExtraPoint(
	t *testing.T) {
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

func TestShouldReturnWeightedAverageOnASliceNumberWithMinimumClassReachedWithTeachersListExtraPoint(
	t *testing.T) {
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

func TestShouldReturnWeightedAverageOnASliceNumberWithMinimumClassNoReachedWithTeachersListExtraPoint(
	t *testing.T) {
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
		s.whenCalculateGradesMinimumReachedAndWeightedAverageAndTeachers(false, "teacher1")

		s.expected = 0
		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

func TestShouldReturnZeroOnEmptySliceMinimumWeightedTeacherListExtraPoint(t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	s.studentGrade = []domain.StudentGrade{}
	s.expected = 0

	s.whenCalculateGradesMinimumReachedAndWeightedAverageAndTeachers(true, "teacher1")
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")

	s.whenCalculateGradesMinimumReachedAndWeightedAverageAndTeachers(false, "teacher2")
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")
}

func TestShouldReturnAnErrorCodeWithWWeightedUnderOneHundredTeacherListExtraPoint(t *testing.T) {
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
		s.whenCalculateGradesMinimumReachedAndWeightedAverageAndTeachers(true, "teacher2")

		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

func TestShouldReturnAnErrorCodeWithWWeightedOverOneHundredTeacherListExtraPoint(t *testing.T) {
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
		s.whenCalculateGradesMinimumReachedAndWeightedAverageAndTeachers(true, "teacher1")

		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

// Upper test repeated with teachers list...

/*-- Feature 5: Refactors */
func TestRefactorShouldReturnWeightedAverageOnASliceNumberWithMinimumClassReachedWithTeachersListNoExtraPoint(
	t *testing.T) {
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
		s.whenCalculateGradesRefactor(true, "teacher2", 1)

		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

func TestRefactorShouldReturnWeightedAverageOnASliceNumberWithMinimumClassReachedWithTeachersListExtraPoint(
	t *testing.T) {
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
		s.whenCalculateGradesRefactor(true, "teacher1", 1)

		s.expected += 1
		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

func TestRefactorShouldReturnWeightedAverageOnASliceNumberWithMinimumClassNoReachedWithTeachersListExtraPoint(
	t *testing.T) {
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
		s.whenCalculateGradesRefactor(false, "teacher1", 1)

		s.expected = 0
		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

func TestRefactorShouldReturnZeroOnEmptySliceMinimumWeightedTeacherListExtraPoint(
	t *testing.T) {
	s := startStudentGradeCalculatorTest(t)

	s.studentGrade = []domain.StudentGrade{}
	s.expected = 0

	s.whenCalculateGradesRefactor(true, "teacher1", 1)
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")

	s.whenCalculateGradesRefactor(false, "teacher2", 1)
	assert.Equal(t, s.expected, s.result, "should return zero on empty slice")
}

func TestRefactorShouldReturnAnErrorCodeWithWWeightedUnderOneHundredTeacherListExtraPoint(
	t *testing.T) {
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
		s.whenCalculateGradesRefactor(true, "teacher2", 1)

		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

func TestRefactorShouldReturnAnErrorCodeWithWWeightedOverOneHundredTeacherListExtraPoint(
	t *testing.T) {
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
		s.whenCalculateGradesRefactor(true, "teacher1", 1)

		assert.Equal(t, s.expected, s.result, "should return average on a slice of grades")
	}
}

/*-- steps ---*/
func startStudentGradeCalculatorTest(t *testing.T) *studentGradeCalculatorTest {
	t.Parallel()

	return &studentGradeCalculatorTest{
		test: t,
	}
}

func (s *studentGradeCalculatorTest) givenASliceOfNumbers(length int) {
	var numbers []domain.Grades
	var sumOfValues domain.Grades

	for i := 0; i < length; i++ {
		numbers = append(numbers, domain.Grades(rand.Int()))
		sumOfValues += numbers[i]
	}

	s.grades = numbers
	s.expected = sumOfValues / domain.Grades(length)
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

	if sumOfWeights > 100 {
		s.expected = -1
	}
	if sumOfWeights < 100 {
		s.expected = -2
	}
}

func (s *studentGradeCalculatorTest) andThereIsAtTeacherExtraPointMap() {
	s.teachers = make(domain.TeacherExtraPoint)
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
	s.result = target.calculateGradesMinimumClassesAndWeightedAverageWithExtraPoint(
		s.studentGrade, hasReachMinimum, teacher)
}

func (s *studentGradeCalculatorTest) whenCalculateGradesRefactor(
	hasReachMinimum bool, teacher string, refactor int) {
	target := NewStudentGradeCalculatorRefactor(s.teachers)
	switch refactor {
	case 1:
		s.result = target.calculateGradesRefactorOne(s.studentGrade, hasReachMinimum, teacher)
	default:
		s.result = 0
	}
}
