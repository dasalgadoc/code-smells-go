package domain

type DateRange struct {
	startDate LocalDateTime
	endDate   LocalDateTime
}

func (d *DateRange) StartDate() LocalDateTime {
	return d.startDate
}

func (d *DateRange) EndDate() LocalDateTime {
	return d.endDate
}

func NewDateRanges(startDate LocalDateTime, endDate LocalDateTime) *DateRange {
	return &DateRange{
		startDate: startDate,
		endDate:   endDate,
	}
}
