package domain

import "errors"

type CourseDuration struct {
	duration int
}

func NewCourseDuration(duration int) (*CourseDuration, error) {
	if duration < 0 {
		return nil, errors.New("duration must be greater than 0")
	}
	return &CourseDuration{duration: duration}, nil
}

func (c *CourseDuration) Value() int {
	return c.duration
}
