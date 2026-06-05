package product_group

import (
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product_group"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/repository/model"
)

func (r *productGroupRepository) FindAll() ([]product_group.ProductGroup, error) {
	rows, err := r.db.Query(`SELECT id, name, created_at, updated_at FROM product_groups`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []product_group.ProductGroup
	for rows.Next() {
		var row model.ProductGroupRow
		if err := rows.Scan(&row.ID, &row.Name, &row.CreatedAt, &row.UpdatedAt); err != nil {
			return nil, err
		}
		groups = append(groups, *row.ToEntity())
	}
	return groups, nil
}
