package migration

import (
	"database/sql"
	"fmt"

	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

type Config struct {
	Driver string
	DSN    string
}

func Run(cfg Config) error {
	db, err := sql.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping db: %w", err)
	}

	if err := Up(db); err != nil {
		return fmt.Errorf("migration up failed: %w", err)
	}

	logger.Info("migration completed", "total", len(GetMigrations()))
	return nil
}

func Status(cfg Config) error {
	db, err := sql.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}
	defer db.Close()

	migrations := GetMigrations()
	logger.Info("migration status", "total_migrations", len(migrations))
	for _, m := range migrations {
		logger.Info("  -", "version", m.Version, "name", m.Name)
	}
	return nil
}

func Drop(cfg Config) error {
	db, err := sql.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}
	defer db.Close()

	if err := DropAll(db); err != nil {
		return fmt.Errorf("drop failed: %w", err)
	}

	logger.Info("all migrations dropped")
	return nil
}

func Version(cfg Config) (int, error) {
	db, err := sql.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	migrations := GetMigrations()
	return len(migrations), nil
}
