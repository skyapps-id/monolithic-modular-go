package product

import (
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product"
)

type ProductHandler struct {
	uc product.ProductUsecase
}

func NewProductHandler(uc product.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}
