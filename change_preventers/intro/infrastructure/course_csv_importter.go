package infrastructure

import (
	"dasalgadoc.com/code_smell_go/change_preventers/intro/domain"
	"encoding/csv"
	"errors"
	"os"
)

type CourseCsvImporter struct {
	location string
}

func NewCourseCsvImporter() domain.Importer {
	return &CourseCsvImporter{
		location: "./data.csv",
	}
}

func (c *CourseCsvImporter) Invoke(course string) (*domain.Table, error) {
	file, err := os.Open(c.location)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		if row[0] == course {
			return &domain.Table{
				Headers: headers,
				Rows:    rows,
			}, nil
		}
	}

	return nil, errors.New("course not found")
}
