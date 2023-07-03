package strategy

import "time"

type MonthLimitGenerator struct{}

func (g *MonthLimitGenerator) Generate(date time.Time) (time.Time, time.Time) {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC),
		time.Date(date.Year(), date.Month()+1, 0, 23, 59, 59, 999, time.UTC)
}

func NewMonthLimitGenerator() *MonthLimitGenerator {
	return &MonthLimitGenerator{}
}
