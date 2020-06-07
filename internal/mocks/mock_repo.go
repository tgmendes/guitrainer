package mocks

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/mock"
)

type Repo struct {
	mock.Mock
}

func (m *Repo) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	callArgs := m.Called(ctx, sql, args)
	return callArgs.Get(0).(pgx.Rows), callArgs.Error(1)
}

func (m *Repo) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	callArgs := m.Called(ctx, sql, args)
	return callArgs.Get(0).(pgx.Row)
}

func (m *Repo) Close(ctx context.Context) error {
	callArgs := m.Called(ctx)
	return callArgs.Error(0)
}

func (m *Repo) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	callArgs := m.Called(ctx, sql, args)
	return callArgs.Get(0).([]byte), callArgs.Error(1)
}
