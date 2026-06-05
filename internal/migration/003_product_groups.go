package migration

import (
	"database/sql"
)

func init() {
	Register(3, "create_product_groups_table", createProductGroupsTable, dropProductGroupsTable)
}

func createProductGroupsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS product_groups (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}

func dropProductGroupsTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS product_groups")
	return err
}
