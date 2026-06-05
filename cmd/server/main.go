package main

import (
	"flag"
	"os"

	_ "github.com/glebarez/sqlite"
	"github.com/skyapps-id/monolithic-modular-go/internal/bootstrap"
	"github.com/skyapps-id/monolithic-modular-go/internal/driver"
	"github.com/skyapps-id/monolithic-modular-go/internal/migration"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users"
	"github.com/skyapps-id/monolithic-modular-go/internal/router"
)

func main() {
	configPath := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "app.db"
	}

	migrateCfg := migration.Config{
		Driver: "sqlite",
		DSN:    dsn,
	}
	if err := migration.Run(migrateCfg); err != nil {
		panic("failed to run migrations: " + err.Error())
	}

	db, err := driver.NewSQLite(dsn)
	if err != nil {
		panic("failed to open db: " + err.Error())
	}
	defer db.Close()

	modules := []router.Module{
		users.NewModule(db),
		products.NewModule(db),
	}

	bootstrap.Run(*configPath, modules)
}
