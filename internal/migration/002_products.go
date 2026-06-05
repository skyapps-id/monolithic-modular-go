package migration

import (
	"database/sql"
)

func init() {
	Register(2, "create_products_table", createProductsTable, dropProductsTable)
}

func createProductsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		price REAL NOT NULL,
		stock INTEGER NOT NULL DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}

func dropProductsTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS products")
	return err
}
