package strategy

import "time"

type LimitDatesRefactor struct {
	Generator Generator
	Type      LimitType
	Gte       time.Time
	Lte       time.Time
}

func (l *LimitDatesRefactor) GenerateDates() {
	now := CurrentDate()
	l.Gte, l.Lte = l.Generator.Generate(now)
}

func NewLimitDatesRefactor(generator Generator, limitType LimitType) *LimitDatesRefactor {
	return &LimitDatesRefactor{
		Generator: generator,
		Type:      limitType,
	}
}
