package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

type Database struct {
	Course CourseDatabase
}

func NewPool(databaseURL string) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		return nil, nil
	}

	return pool, nil
}

func NewDatabase(databaseURL string) (*Database, error) {
	pool, err := NewPool(databaseURL)

	if err != nil {
		return nil, err
	}

	return &Database{
		Course: CourseDatabase{Pool: pool},
	}, nil
}
