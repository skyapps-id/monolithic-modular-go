package migration

import (
	"database/sql"
	"fmt"
)

type Migration struct {
	Version uint
	Name    string
	Up      func(*sql.DB) error
	Down    func(*sql.DB) error
}

var migrations []Migration

func Register(v uint, name string, up, down func(*sql.DB) error) {
	migrations = append(migrations, Migration{Version: v, Name: name, Up: up, Down: down})
}

func GetMigrations() []Migration {
	return migrations
}

func Up(db *sql.DB) error {
	for _, m := range migrations {
		if err := m.Up(db); err != nil {
			return fmt.Errorf("migration %d (%s) failed: %w", m.Version, m.Name, err)
		}
	}
	return nil
}

func Down(db *sql.DB) error {
	for i := len(migrations) - 1; i >= 0; i-- {
		m := migrations[i]
		if err := m.Down(db); err != nil {
			return fmt.Errorf("rollback %d (%s) failed: %w", m.Version, m.Name, err)
		}
	}
	return nil
}

func DropAll(db *sql.DB) error {
	for i := len(migrations) - 1; i >= 0; i-- {
		m := migrations[i]
		if err := m.Down(db); err != nil {
			return fmt.Errorf("drop %d (%s) failed: %w", m.Version, m.Name, err)
		}
	}
	return nil
}
