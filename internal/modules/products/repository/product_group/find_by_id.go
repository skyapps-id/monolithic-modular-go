package product_group

import (
	"database/sql"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product_group"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/repository/model"
)

func (r *productGroupRepository) FindByID(id string) (*product_group.ProductGroup, error) {
	var row model.ProductGroupRow
	err := r.db.QueryRow(
		`SELECT id, name, created_at, updated_at FROM product_groups WHERE id = ?`, id,
	).Scan(&row.ID, &row.Name, &row.CreatedAt, &row.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return row.ToEntity(), nil
}

var _ = (*sql.DB)(nil)
