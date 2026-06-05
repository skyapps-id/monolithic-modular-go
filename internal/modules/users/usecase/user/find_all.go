package user

import (
	"context"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/domain/user"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

func (u *userUsecase) FindAll(ctx context.Context) ([]user.User, error) {
	logger.DebugCtx(ctx, "user.usecase.find_all.start")

	entities, err := u.repo.FindAll()
	if err != nil {
		logger.ErrorCtx(ctx, "user.usecase.find_all.error", "error", err.Error())
		return nil, apperror.New(apperror.InternalError, "failed to fetch users")
	}

	logger.InfoCtx(ctx, "user.usecase.find_all.success", "count", len(entities))
	return entities, nil
}
