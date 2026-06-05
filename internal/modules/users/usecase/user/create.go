package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/domain/user"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

func (u *userUsecase) Create(ctx context.Context, input *user.CreateUserInput) (*user.User, error) {
	logger.DebugCtx(ctx, "user.usecase.create.start", "email", input.Email)

	entity := &user.User{
		ID:    uuid.New().String(),
		Name:  input.Name,
		Email: input.Email,
	}

	now := time.Now()
	entity.CreatedAt = now
	entity.UpdatedAt = now

	if err := u.repo.Save(entity); err != nil {
		logger.ErrorCtx(ctx, "user.usecase.create.error", "error", err.Error())
		return nil, apperror.New(apperror.InternalError, "failed to save user")
	}

	logger.InfoCtx(ctx, "user.usecase.create.success", "user_id", entity.ID)
	return entity, nil
}
