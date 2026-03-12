package db

import (
	"context"
	"ems/config"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgDb struct {
	Pool *pgxpool.Pool
}

func ConnectDb(cfg *config.ConfigStruct) (*PgDb, error) {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable pool_max_conns=10",
		cfg.DBIP,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &PgDb{Pool: pool}, nil
}
