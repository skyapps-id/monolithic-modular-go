package product_group

import (
	"context"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product_group"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

func (u *productGroupUsecase) FindAll(ctx context.Context) ([]product_group.ProductGroup, error) {
	logger.DebugCtx(ctx, "product_group.usecase.find_all.start")

	entities, err := u.repo.FindAll()
	if err != nil {
		logger.ErrorCtx(ctx, "product_group.usecase.find_all.error", "error", err.Error())
		return nil, apperror.New(apperror.InternalError, "failed to fetch product groups")
	}

	logger.InfoCtx(ctx, "product_group.usecase.find_all.success", "count", len(entities))
	return entities, nil
}
