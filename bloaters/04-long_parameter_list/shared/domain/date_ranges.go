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

func (d *DateRange) HasStarted(date LocalDateTime) bool {
	return d.startDate.IsAfter(date)
}

func (d *DateRange) IsBetween(date LocalDateTime) bool {
	return d.startDate.IsBefore(date) && d.endDate.IsAfter(date)
}

func NewDateRanges(startDate LocalDateTime, endDate LocalDateTime) *DateRange {
	return &DateRange{
		startDate: startDate,
		endDate:   endDate,
	}
}
