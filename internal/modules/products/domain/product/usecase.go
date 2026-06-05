package product

import "context"

type CreateProductInput struct {
	Name  string
	Price float64
	Stock int
}

type ProductUsecase interface {
	Create(ctx context.Context, input *CreateProductInput) (*Product, error)
	FindByID(ctx context.Context, id string) (*Product, error)
	FindAll(ctx context.Context) ([]Product, error)
}
