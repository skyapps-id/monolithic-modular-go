package product

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

func (u *productUsecase) Create(ctx context.Context, input *product.CreateProductInput) (*product.Product, error) {
	logger.DebugCtx(ctx, "product.usecase.create.start", "name", input.Name)

	entity := &product.Product{
		ID:    uuid.New().String(),
		Name:  input.Name,
		Price: input.Price,
		Stock: input.Stock,
	}

	now := time.Now()
	entity.CreatedAt = now
	entity.UpdatedAt = now

	if err := u.repo.Save(entity); err != nil {
		logger.ErrorCtx(ctx, "product.usecase.create.error", "error", err.Error())
		return nil, apperror.New(apperror.InternalError, "failed to save product")
	}

	logger.InfoCtx(ctx, "product.usecase.create.success", "product_id", entity.ID)
	return entity, nil
}
