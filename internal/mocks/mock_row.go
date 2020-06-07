package mocks

import (
	"github.com/stretchr/testify/mock"
)

type Row struct {
	mock.Mock
}

func (m *Row) Scan(dest ...interface{}) error {
	callArgs := m.Called(dest)

	return callArgs.Error(0)
}
