package application

import (
	"dasalgadoc.com/code_smell_go/change_preventers/intro/infrastructure"
	"dasalgadoc.com/code_smell_go/change_preventers/intro/infrastructure/dto"
	"encoding/json"
	"strconv"
)

type StepsCalculator struct {
	csvImporter infrastructure.CsvImporter
}

func NewStepsCalculator(csvImporter infrastructure.CsvImporter) *StepsCalculator {
	return &StepsCalculator{
		csvImporter: csvImporter,
	}
}

func (s *StepsCalculator) Get(courseId string) (*string, error) {
	table, err := s.csvImporter.Get(courseId)
	if err != nil {
		return nil, err
	}

	result := "["
	for _, row := range table.Rows {
		stepType := row[0]
		duration, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, err
		}

		points, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, err
		}

		if stepType == "video" {
			duration = duration * 1.1
		}

		if stepType == "quiz" {
			duration = duration * 1.1
		}

		if stepType != "video" && stepType != "quiz" {
			continue
		}

		if stepType == "video" {
			points = points * 1.1 * 100
		}

		if stepType == "quiz" {
			points = points * 0.5 * 10
		}

		var data dto.Course
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		result = result + string(jsonData) + ","

	}
	result = result + "]"
	return &result, nil
}
