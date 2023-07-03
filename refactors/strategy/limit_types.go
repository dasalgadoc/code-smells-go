package strategy

type LimitType string

const (
	DayLimitType   LimitType = "day"
	MonthLimitType LimitType = "month"
	YearLimitType  LimitType = "year"
)
