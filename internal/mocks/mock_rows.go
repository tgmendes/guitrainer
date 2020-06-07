package mocks

import (
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/mock"
)

type Rows struct {
	pgx.Rows
	mock.Mock
}

func (m *Rows) Next() bool {
	callArgs := m.Called()

	return callArgs.Bool(0)
}

func (m *Rows) Scan(dest ...interface{}) error {
	callArgs := m.Called(dest)

	return callArgs.Error(0)
}
