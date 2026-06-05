package product_group

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product_group"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

func (u *productGroupUsecase) Create(ctx context.Context, input *product_group.CreateProductGroupInput) (*product_group.ProductGroup, error) {
	logger.DebugCtx(ctx, "product_group.usecase.create.start", "name", input.Name)

	entity := &product_group.ProductGroup{
		ID:   uuid.New().String(),
		Name: input.Name,
	}

	now := time.Now()
	entity.CreatedAt = now
	entity.UpdatedAt = now

	if err := u.repo.Save(entity); err != nil {
		logger.ErrorCtx(ctx, "product_group.usecase.create.error", "error", err.Error())
		return nil, apperror.New(apperror.InternalError, "failed to save product group")
	}

	logger.InfoCtx(ctx, "product_group.usecase.create.success", "group_id", entity.ID)
	return entity, nil
}
