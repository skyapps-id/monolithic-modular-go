package product_group

import (
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product_group"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/repository/model"
)

func (r *productGroupRepository) Save(entity *product_group.ProductGroup) error {
	row := model.ProductGroupFromEntity(entity)
	_, err := r.db.Exec(
		`INSERT INTO product_groups (id, name, created_at, updated_at) VALUES (?, ?, ?, ?)`,
		row.ID, row.Name, row.CreatedAt, row.UpdatedAt,
	)
	return err
}
