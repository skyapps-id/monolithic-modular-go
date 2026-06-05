package user

import (
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/domain/user"
)

type UserHandler struct {
	uc user.UserUsecase
}

func NewUserHandler(uc user.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}
