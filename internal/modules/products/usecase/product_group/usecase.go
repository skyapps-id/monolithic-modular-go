package product_group

import (
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product_group"
)

var _ product_group.ProductGroupUsecase = (*productGroupUsecase)(nil)

type productGroupUsecase struct {
	repo product_group.ProductGroupRepository
}

func NewProductGroupUsecase(repo product_group.ProductGroupRepository) product_group.ProductGroupUsecase {
	return &productGroupUsecase{repo: repo}
}
