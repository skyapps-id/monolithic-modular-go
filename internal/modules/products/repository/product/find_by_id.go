package product

import (
	"database/sql"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/repository/model"
)

func (r *productRepository) FindByID(id string) (*product.Product, error) {
	var row model.ProductRow
	err := r.db.QueryRow(
		`SELECT id, name, price, stock, created_at, updated_at FROM products WHERE id = ?`, id,
	).Scan(&row.ID, &row.Name, &row.Price, &row.Stock, &row.CreatedAt, &row.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return row.ToEntity(), nil
}

var _ = (*sql.DB)(nil)
