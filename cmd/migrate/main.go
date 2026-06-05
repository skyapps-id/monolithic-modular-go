package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/glebarez/sqlite"
	"github.com/skyapps-id/monolithic-modular-go/internal/migration"
)

func main() {
	migrateUp := flag.Bool("up", false, "run migrations up")
	migrateDown := flag.Bool("down", false, "run migrations down")
	drop := flag.Bool("drop", false, "drop all tables")
	status := flag.Bool("status", false, "show migration status")
	flag.Parse()

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "app.db"
	}

	cfg := migration.Config{
		Driver: "sqlite",
		DSN:    dsn,
	}

	switch {
	case *drop:
		if err := migration.Drop(cfg); err != nil {
			log.Fatalf("drop failed: %v", err)
		}
		log.Println("all tables dropped")

	case *status:
		if err := migration.Status(cfg); err != nil {
			log.Fatalf("status failed: %v", err)
		}

	case *migrateDown:
		db, err := sql.Open(cfg.Driver, cfg.DSN)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		if err := migration.Down(db); err != nil {
			log.Fatalf("migration down failed: %v", err)
		}
		log.Println("migrated down")

	case *migrateUp:
		if err := migration.Run(cfg); err != nil {
			log.Fatalf("migration up failed: %v", err)
		}
		log.Println("migration completed")

	default:
		fmt.Println("Usage: migrate -up | -down | -drop | -status")
	}
}
