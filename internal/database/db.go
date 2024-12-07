package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"todo/internal/configs"
)

type Db struct {
	Pool *pgxpool.Pool
}

func NewDB(ctx context.Context, conf *configs.Config) *Db {
	cfg, err := pgxpool.ParseConfig(conf.Db.Dsn)
	if err != nil {
		return nil
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil
	}

	return &Db{Pool: pool}
}
