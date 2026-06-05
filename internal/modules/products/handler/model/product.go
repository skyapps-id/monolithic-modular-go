package model

import (
	"time"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product"
)

type CreateProductRequest struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,gt=0"`
	Stock int     `json:"stock"`
}

type ProductResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *CreateProductRequest) ToInput() *product.CreateProductInput {
	return &product.CreateProductInput{
		Name:  r.Name,
		Price: r.Price,
		Stock: r.Stock,
	}
}

func ProductToResponse(p *product.Product) *ProductResponse {
	return &ProductResponse{
		ID:        p.ID,
		Name:      p.Name,
		Price:     p.Price,
		Stock:     p.Stock,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ProductsToResponse(products []product.Product) []ProductResponse {
	res := make([]ProductResponse, len(products))
	for i, p := range products {
		res[i] = *ProductToResponse(&p)
	}
	return res
}
