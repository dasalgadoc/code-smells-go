package strategy

import "time"

type YearLimitGenerator struct{}

func (g *YearLimitGenerator) Generate(date time.Time) (time.Time, time.Time) {
	return time.Date(date.Year(), 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(date.Year(), 12, 31, 23, 59, 59, 999, time.UTC)
}

func NewYearLimitGenerator() *YearLimitGenerator {
	return &YearLimitGenerator{}
}
