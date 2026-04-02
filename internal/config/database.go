package config

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func ConnectionDB() (*sql.DB, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return nil, errors.New("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, errors.New("Failed to connect to database: " + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		closeErr := db.Close()
		if closeErr != nil {
			return nil, errors.New("ping database: " + err.Error() + " (close connection: " + closeErr.Error() + ")")
		}
		return nil, errors.New("ping database: " + err.Error())

	}

	return db, nil
}
