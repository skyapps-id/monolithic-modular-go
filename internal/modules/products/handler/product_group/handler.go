package product_group

import (
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product_group"
)

type ProductGroupHandler struct {
	uc product_group.ProductGroupUsecase
}

func NewProductGroupHandler(uc product_group.ProductGroupUsecase) *ProductGroupHandler {
	return &ProductGroupHandler{uc: uc}
}
