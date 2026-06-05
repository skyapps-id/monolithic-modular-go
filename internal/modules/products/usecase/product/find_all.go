package product

import (
	"context"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

func (u *productUsecase) FindAll(ctx context.Context) ([]product.Product, error) {
	logger.DebugCtx(ctx, "product.usecase.find_all.start")

	entities, err := u.repo.FindAll()
	if err != nil {
		logger.ErrorCtx(ctx, "product.usecase.find_all.error", "error", err.Error())
		return nil, apperror.New(apperror.InternalError, "failed to fetch products")
	}

	logger.InfoCtx(ctx, "product.usecase.find_all.success", "count", len(entities))
	return entities, nil
}
