package test

import (
	"dasalgadoc.com/code_smell_go/change_preventers/intro/domain"
	"github.com/stretchr/testify/mock"
)

type ImporterMock struct {
	mock.Mock
}

func (m *ImporterMock) Invoke(courseId string) (*domain.Table, error) {
	args := m.Called(courseId)
	return args.Get(0).(*domain.Table), args.Error(1)
}
