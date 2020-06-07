package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Repository interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Close(ctx context.Context) error
}

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func Open(ctx context.Context, cfg Config) (*pgx.Conn, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	conn, err := pgx.Connect(ctx, dbURL)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
