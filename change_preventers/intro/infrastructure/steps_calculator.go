package infrastructure

import (
	"dasalgadoc.com/code_smell_go/change_preventers/intro/infrastructure/dto"
	"encoding/json"
	"strconv"
)

type StepsCalculatorController struct {
	csvImporter CsvImporter
}

func NewStepsCalculatorController(csvImporter CsvImporter) *StepsCalculatorController {
	return &StepsCalculatorController{
		csvImporter: csvImporter,
	}
}

func (s *StepsCalculatorController) Get(courseId string) (*string, error) {
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
