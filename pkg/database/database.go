package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/riskibarqy/ihsan-test/internal/config"
)

func ConnectPostgres(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.GetDatabaseDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	return db, nil
}
