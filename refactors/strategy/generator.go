package strategy

import "time"

type Generator interface {
	Generate(date time.Time) (time.Time, time.Time)
}
