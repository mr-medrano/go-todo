package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseConfig struct {
	Username string
	Password string
	Hostname string
	Port     int
	DBName   string
}

func (db DatabaseConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		db.Username, db.Password, db.Hostname, db.Port, db.DBName)
}

func NewDBPool(dbConfig DatabaseConfig) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dbConfig.DSN())
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
