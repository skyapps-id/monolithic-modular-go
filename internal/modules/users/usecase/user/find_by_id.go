package user

import (
	"context"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/domain/user"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

func (u *userUsecase) FindByID(ctx context.Context, id string) (*user.User, error) {
	logger.DebugCtx(ctx, "user.usecase.find_by_id.start", "id", id)

	entity, err := u.repo.FindByID(id)
	if err != nil {
		logger.ErrorCtx(ctx, "user.usecase.find_by_id.error", "error", err.Error())
		return nil, err
	}
	if entity == nil {
		logger.WarnCtx(ctx, "user.usecase.find_by_id.not_found", "id", id)
		return nil, apperror.NotFoundError("user")
	}

	logger.InfoCtx(ctx, "user.usecase.find_by_id.success", "id", id)
	return entity, nil
}
