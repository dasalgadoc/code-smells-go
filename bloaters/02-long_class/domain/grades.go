package domain

import (
	"errors"
	"fmt"
)

type Grades struct {
	Value []Grade
}

func (g *Grades) SumWeightedGrades() NumericGrade {
	var weightedGradesSum NumericGrade
	for _, grade := range g.Value {
		weightedGradesSum += grade.CalculateGrade()
	}
	return weightedGradesSum
}

func (g *Grades) CalculateAverage() NumericGrade {
	return g.SumWeightedGrades() / NumericGrade(len(g.Value))
}

func NewGrades(values []NumericGrade, weights []int) (*Grades, error) {
	if len(values) != len(weights) {
		return nil, errors.New("the number of values and weights are not the same")
	}

	err := checkErrorOnWeights(weights)
	if err != nil {
		return nil, err
	}

	err = checkErrorOnValues(values)
	if err != nil {
		return nil, err
	}

	var Grades Grades
	for i, value := range values {
		Grades.Value = append(Grades.Value, Grade{
			Value:  value,
			Weight: weights[i],
		})
	}

	return &Grades, nil
}

func checkErrorOnWeights(weights []int) error {
	var sumOfWeights int
	for _, weight := range weights {
		sumOfWeights += weight
	}

	if sumOfWeights > 100 {
		return errors.New("the average is over-weighed")
	}
	if sumOfWeights < 100 {
		return errors.New("the average is under-weighed")
	}
	return nil
}

func checkErrorOnValues(values []NumericGrade) error {
	for _, value := range values {
		if value < 0 || value > 10 {
			errorStr := fmt.Sprintf("the value %d is not between 0 and 100", value)
			return errors.New(errorStr)
		}
	}
	return nil
}
