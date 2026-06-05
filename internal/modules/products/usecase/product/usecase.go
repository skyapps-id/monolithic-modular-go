package product

import (
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product"
)

var _ product.ProductUsecase = (*productUsecase)(nil)

type productUsecase struct {
	repo product.ProductRepository
}

func NewProductUsecase(repo product.ProductRepository) product.ProductUsecase {
	return &productUsecase{repo: repo}
}
