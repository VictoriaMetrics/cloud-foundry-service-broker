package db

import (
	"errors"
	"github.com/VictoriaMetrics/cloud-foundry-service-broker/lib/logger"
	"github.com/golang-migrate/migrate/v4"
	"os"
	"strings"

	_ "github.com/golang-migrate/migrate/v4/database/pgx" // Migrations driver for pgx
	_ "github.com/golang-migrate/migrate/v4/source/file"  // Migrations driver for file
)

func migrationsDirExists() bool {
	if _, err := os.Stat("migrations"); os.IsNotExist(err) {
		return false
	}

	return true
}

// Migrate runs database migrations if migrations files are present
func Migrate() {
	if !migrationsDirExists() {
		logger.DefaultLogger.Warn("migrations directory not found, skipped migrations")
		return
	}

	newDbURL := strings.Replace(*dbURL, "postgres://", "pgx://", -1)
	m, err := migrate.New("file://migrations", newDbURL)
	if err != nil {
		logger.DefaultLogger.Fatalw("failed to create migrator", "error", err)
	}

	if err := m.Up(); errors.Is(err, migrate.ErrNoChange) {
		logger.DefaultLogger.Debug("no migration needed")
	} else if err != nil {
		logger.DefaultLogger.Fatalw("failed to migrate", "error", err)
	}

}
