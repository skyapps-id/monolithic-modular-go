package product

import (
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/repository/model"
)

func (r *productRepository) Save(entity *product.Product) error {
	row := model.ProductFromEntity(entity)
	_, err := r.db.Exec(
		`INSERT INTO products (id, name, price, stock, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`,
		row.ID, row.Name, row.Price, row.Stock, row.CreatedAt, row.UpdatedAt,
	)
	return err
}
