package domain

import "errors"

type CourseName struct {
	name string
}

func NewCourseName(name string) (*CourseName, error) {
	if len(name) < 1 {
		return nil, errors.New("name cannot be empty")
	}
	return &CourseName{name: name}, nil
}

func (c *CourseName) Value() string {
	return c.name
}
