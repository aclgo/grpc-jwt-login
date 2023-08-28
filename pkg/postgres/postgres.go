package postgres

import (
	"context"
	"fmt"

	"github.com/aclgo/grpc-jwt/config"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect(cfg *config.Config) (*sqlx.DB, error) {
	conn, err := sqlx.Connect(cfg.DatabaseDriver, cfg.DatabaseUrl)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Connect %v", err)
	}

	if err := conn.PingContext(context.Background()); err != nil {
		return nil, fmt.Errorf("PingContext %v", err)
	}

	return conn, nil
}
