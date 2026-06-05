package model

import (
	"time"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product_group"
)

type CreateProductGroupRequest struct {
	Name string `json:"name" validate:"required"`
}

type ProductGroupResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *CreateProductGroupRequest) ToInput() *product_group.CreateProductGroupInput {
	return &product_group.CreateProductGroupInput{
		Name: r.Name,
	}
}

func GroupToResponse(g *product_group.ProductGroup) *ProductGroupResponse {
	return &ProductGroupResponse{
		ID:        g.ID,
		Name:      g.Name,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}

func GroupsToResponse(groups []product_group.ProductGroup) []ProductGroupResponse {
	res := make([]ProductGroupResponse, len(groups))
	for i, g := range groups {
		res[i] = *GroupToResponse(&g)
	}
	return res
}
