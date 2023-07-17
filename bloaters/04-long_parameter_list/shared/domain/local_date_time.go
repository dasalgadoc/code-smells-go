package domain

import "time"

type LocalDateTime struct {
	value time.Time
}

func (l *LocalDateTime) Value() time.Time {
	return l.value
}

func (l *LocalDateTime) IsBefore(other LocalDateTime) bool {
	return l.value.Before(other.value)
}

func (l *LocalDateTime) IsAfter(other LocalDateTime) bool {
	return l.value.After(other.value)
}

func (l *LocalDateTime) IsBetween(start, end LocalDateTime) bool {
	return l.IsAfter(start) && l.IsBefore(end)
}

func NewLocalDateTime() *LocalDateTime {
	return &LocalDateTime{value: time.Now()}
}

func NewLocalDateTimeFromTime(t time.Time) *LocalDateTime {
	return &LocalDateTime{value: t}
}
