package bootstrap

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"restaurant/internal/config"
)

func InitSqlxDB(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", formatConnect(cfg))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func formatConnect(cfg config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresSQL.Username, cfg.PostgresSQL.Password, cfg.PostgresSQL.Host, cfg.PostgresSQL.Port, cfg.PostgresSQL.Database,
	)
}
