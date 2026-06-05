package product_group

import (
	"context"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product_group"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

func (u *productGroupUsecase) FindByID(ctx context.Context, id string) (*product_group.ProductGroup, error) {
	logger.DebugCtx(ctx, "product_group.usecase.find_by_id.start", "id", id)

	entity, err := u.repo.FindByID(id)
	if err != nil {
		logger.ErrorCtx(ctx, "product_group.usecase.find_by_id.error", "error", err.Error())
		return nil, err
	}
	if entity == nil {
		logger.WarnCtx(ctx, "product_group.usecase.find_by_id.not_found", "id", id)
		return nil, apperror.NotFoundError("product group")
	}

	logger.InfoCtx(ctx, "product_group.usecase.find_by_id.success", "id", id)
	return entity, nil
}
