package product_group

import "context"

type CreateProductGroupInput struct {
	Name string
}

type ProductGroupUsecase interface {
	Create(ctx context.Context, input *CreateProductGroupInput) (*ProductGroup, error)
	FindByID(ctx context.Context, id string) (*ProductGroup, error)
	FindAll(ctx context.Context) ([]ProductGroup, error)
}
