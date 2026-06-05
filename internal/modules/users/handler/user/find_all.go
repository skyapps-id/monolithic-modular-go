package user

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/handler/model"
	"github.com/skyapps-id/monolithic-modular-go/pkg/response"
)

// FindAll godoc
//
//	@Summary		List users
//	@Description	Get all users
//	@Tags			users
//	@Produce		json
//	@Success		200		{object}	response.SuccessResponse
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/api/v1/users [get]
func (h *UserHandler) FindAll(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := h.uc.FindAll(ctx)
	if err != nil {
		return err
	}

	return response.OK(c, model.UsersToResponse(users))
}
