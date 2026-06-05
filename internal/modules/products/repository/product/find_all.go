package product

import (
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/repository/model"
)

func (r *productRepository) FindAll() ([]product.Product, error) {
	rows, err := r.db.Query(`SELECT id, name, price, stock, created_at, updated_at FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []product.Product
	for rows.Next() {
		var row model.ProductRow
		if err := rows.Scan(&row.ID, &row.Name, &row.Price, &row.Stock, &row.CreatedAt, &row.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, *row.ToEntity())
	}
	return products, nil
}
