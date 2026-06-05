package model

import (
	"time"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product_group"
)

type ProductGroupRow struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (r *ProductGroupRow) ToEntity() *product_group.ProductGroup {
	return &product_group.ProductGroup{
		ID:        r.ID,
		Name:      r.Name,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func ProductGroupFromEntity(g *product_group.ProductGroup) *ProductGroupRow {
	return &ProductGroupRow{
		ID:        g.ID,
		Name:      g.Name,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}
