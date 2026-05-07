package config

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Database() (*pgxpool.Pool, error) {
	dbURL := GetDatabaseURL()
	if dbURL == "" {
		return nil, fmt.Errorf("database connection string is empty: check your DATABASE_URL environment variable")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	if err := db.Ping(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database (host may be unreachable or credentials invalid): %w", err)
	}

	return db, nil
}