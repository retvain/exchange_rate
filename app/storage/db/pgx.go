package db

import "github.com/jackc/pgx/v4/pgxpool"

type pgxStorage struct {
	db pgxpool.Pool
}

func NewPgxStorage(dsn string) (*pgxStorage, error) {
	cfg := pgxpool.ParseConfig(dsn)
	return &pgxStorage{
		db: nil,
	}
}
