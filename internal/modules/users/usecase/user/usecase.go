package user

import (
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/domain/user"
)

var _ user.UserUsecase = (*userUsecase)(nil)

type userUsecase struct {
	repo user.UserRepository
}

func NewUserUsecase(repo user.UserRepository) user.UserUsecase {
	return &userUsecase{repo: repo}
}
