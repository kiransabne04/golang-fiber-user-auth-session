package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// This file manages the database connection pool setup, retry logic, and returns a connected pool. It uses the configuration struct for the DSN.
var (
	maxRetryAttempts = 10
	retryDelay       = 2 * time.Second
	DbTimeout        = 5 * time.Second
)

// connectToDB initializes connection pool to postgresql database
func ConnectToDB(dburl string) (*pgxpool.Pool, error) {
	for attempt := 1; attempt <= maxRetryAttempts; attempt++ {
		pool, err := openDBPool(dburl)
		if err == nil {
			log.Println("Connected to Postgresql database server")
			return pool, nil
		}

		log.Printf("Failed to connect Postgresql (attempt : %d): %v", attempt, err)
		time.Sleep(retryDelay)
	}
	return nil, fmt.Errorf("unable to connect postgres database instance after %d attempts", maxRetryAttempts)
}

func openDBPool(dsn string) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("error parsing db config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, fmt.Errorf("error creating DB pool: %w", err)
	}
	return pool, nil
}

// pool.QueryRow.Scan - When you have 1 row

// pool.Query - When you have lots of rows and want to op on each row, you would then iterate through pgx.Rows one by one

// pool.Exec - Update and deletes