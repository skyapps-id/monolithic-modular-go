package main

import (
	"database/sql"
	"flag"

	"github.com/skyapps-id/monolithic-modular-go/internal/bootstrap"
	"github.com/skyapps-id/monolithic-modular-go/internal/driver"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users"
	"github.com/skyapps-id/monolithic-modular-go/internal/router"
)

func main() {
	configPath := flag.String("config", "config.inventory.yml", "path to config file")
	flag.Parse()

	db := openDB()
	defer db.Close()

	modules := []router.Module{
		users.NewModule(db),
		products.NewModule(db),
	}

	bootstrap.Run(*configPath, modules)
}

func openDB() *sql.DB {
	db, err := driver.NewSQLite("")
	if err != nil {
		panic(err)
	}
	return db
}
