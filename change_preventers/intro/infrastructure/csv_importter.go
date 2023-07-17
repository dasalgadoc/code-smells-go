package infrastructure

import (
	"dasalgadoc.com/code_smell_go/change_preventers/intro/infrastructure/dto"
	"encoding/csv"
	"os"
)

type CsvImporter struct{}

func NewCsvImporter() *CsvImporter {
	return &CsvImporter{}
}

func (c *CsvImporter) Get(location string) (*dto.Table, error) {
	file, err := os.Open(location)
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

	return &dto.Table{
		Headers: headers,
		Rows:    rows,
	}, nil
}
