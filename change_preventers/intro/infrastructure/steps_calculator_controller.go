package infrastructure

import (
	"dasalgadoc.com/code_smell_go/change_preventers/intro/domain"
	"dasalgadoc.com/code_smell_go/change_preventers/intro/infrastructure/dto"
	"encoding/json"
	"strconv"
)

const (
	VIDEO_DURATION_PAUSES_MULTIPLIER   = 1.1
	QUIZ_TIME_PER_QUESTIONS_MULTIPLIER = 0.5
)

type StepsCalculatorController struct {
	importer domain.Importer
}

func NewStepsCalculatorController(csvImporter domain.Importer) *StepsCalculatorController {
	return &StepsCalculatorController{
		importer: csvImporter,
	}
}

func (s *StepsCalculatorController) Get(courseId string) (*string, error) {
	table, err := s.importer.Invoke(courseId)
	if err != nil {
		return nil, err
	}

	result := "["
	if table != nil {
		for _, row := range table.Rows {
			var stepDuration float64
			var points float64
			stepType := row[1]

			quizQuestions, err := strconv.ParseFloat(row[2], 64)
			if err != nil {
				return nil, err
			}

			videoDuration, err := strconv.ParseFloat(row[3], 64)
			if err != nil {
				return nil, err
			}

			if stepType == "video" {
				stepDuration = videoDuration * VIDEO_DURATION_PAUSES_MULTIPLIER
			}

			if stepType == "quiz" {
				stepDuration = quizQuestions * QUIZ_TIME_PER_QUESTIONS_MULTIPLIER
			}

			if stepType != "video" && stepType != "quiz" {
				continue
			}

			if stepType == "video" {
				points = videoDuration * VIDEO_DURATION_PAUSES_MULTIPLIER * 100
			}

			if stepType == "quiz" {
				points = quizQuestions * QUIZ_TIME_PER_QUESTIONS_MULTIPLIER * 10
			}

			data := dto.CourseJsonDto{
				Id:         row[0],
				CourseType: stepType,
				Duration:   stepDuration,
				Points:     points,
			}

			jsonData, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}

			result = result + string(jsonData) + ","

		}
	}
	if len(result) > 1 {
		result = result[:len(result)-1]
	}
	result = result + "]"
	return &result, nil
}
