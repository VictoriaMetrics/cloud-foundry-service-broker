package db

import (
	"context"
	"github.com/VictoriaMetrics/cloud-foundry-service-broker/lib/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	defaultPool *pgxpool.Pool
)

// Connect Creates connection pool to the database
func Connect() {
	pool, err := pgxpool.Connect(context.Background(), *dbURL)
	if err != nil {
		logger.DefaultLogger.Fatalw("failed to connect to database", "error", err)
	}
	defaultPool = pool
	logger.DefaultLogger.Debug("successfully connected to database")
}

// Close closes the database connection pool
func Close() {
	if defaultPool != nil {
		defaultPool.Close()
	}
}
