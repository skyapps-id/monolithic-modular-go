package model

import (
	"time"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product"
)

type ProductRow struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name"`
	Price     float64   `gorm:"column:price"`
	Stock     int       `gorm:"column:stock"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (r *ProductRow) ToEntity() *product.Product {
	return &product.Product{
		ID:        r.ID,
		Name:      r.Name,
		Price:     r.Price,
		Stock:     r.Stock,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func ProductFromEntity(p *product.Product) *ProductRow {
	return &ProductRow{
		ID:        p.ID,
		Name:      p.Name,
		Price:     p.Price,
		Stock:     p.Stock,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
