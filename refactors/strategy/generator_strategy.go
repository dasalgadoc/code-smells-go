package strategy

func NewGeneratorStrategy(limitType LimitType) Generator {
	return builderStrategyMatcher(limitType)
}

func builderStrategyMatcher(limitType LimitType) Generator {
	return getEnabledDatesRangeGeneratorStrategies()[limitType]
}

func getEnabledDatesRangeGeneratorStrategies() map[LimitType]Generator {
	year := NewYearLimitGenerator()
	month := NewMonthLimitGenerator()
	day := NewDayLimitGenerator()
	generators := map[LimitType]Generator{
		DayLimitType:   day,
		MonthLimitType: month,
		YearLimitType:  year,
	}

	return generators
}
