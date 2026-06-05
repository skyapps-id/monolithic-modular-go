package product

import (
	"context"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

func (u *productUsecase) FindByID(ctx context.Context, id string) (*product.Product, error) {
	logger.DebugCtx(ctx, "product.usecase.find_by_id.start", "id", id)

	entity, err := u.repo.FindByID(id)
	if err != nil {
		logger.ErrorCtx(ctx, "product.usecase.find_by_id.error", "error", err.Error())
		return nil, err
	}
	if entity == nil {
		logger.WarnCtx(ctx, "product.usecase.find_by_id.not_found", "id", id)
		return nil, apperror.NotFoundError("product")
	}

	logger.InfoCtx(ctx, "product.usecase.find_by_id.success", "id", id)
	return entity, nil
}
