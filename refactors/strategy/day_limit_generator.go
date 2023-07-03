package strategy

import "time"

type DayLimitGenerator struct{}

func (g *DayLimitGenerator) Generate(date time.Time) (time.Time, time.Time) {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC),
		time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999, time.UTC)
}

func NewDayLimitGenerator() *DayLimitGenerator {
	return &DayLimitGenerator{}
}
