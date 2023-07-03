package strategy

import "time"

var CurrentDate = func() time.Time {
	return time.Now()
}

type LimitDates struct {
	Type LimitType
	Gte  time.Time
	Lte  time.Time
}

func (l *LimitDates) GenerateDates() {
	nov := CurrentDate()
	if l.Type == DayLimitType {
		l.Gte = time.Date(nov.Year(), nov.Month(), nov.Day(), 0, 0, 0, 0, time.UTC)
		l.Lte = time.Date(nov.Year(), nov.Month(), nov.Day(), 23, 59, 59, 999, time.UTC)
	}
	if l.Type == MonthLimitType {
		l.Gte = time.Date(nov.Year(), nov.Month(), 1, 0, 0, 0, 0, time.UTC)
		l.Lte = time.Date(nov.Year(), nov.Month()+1, 0, 23, 59, 59, 999, time.UTC)
	}
	if l.Type == YearLimitType {
		l.Gte = time.Date(nov.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
		l.Lte = time.Date(nov.Year(), 12, 31, 23, 59, 59, 999, time.UTC)
	}
}
