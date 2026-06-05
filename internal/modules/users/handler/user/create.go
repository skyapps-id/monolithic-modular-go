package user

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/handler/model"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/response"
)

// Create godoc
//
//	@Summary		Create user
//	@Description	Create new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		object	true	"User data"
//	@Success		201		{object}	response.SuccessResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Router			/api/v1/users [post]
func (h *UserHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var req model.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return apperror.ValidationError(err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	created, err := h.uc.Create(ctx, req.ToInput())
	if err != nil {
		return err
	}

	return response.Created(c, model.UserToResponse(created))
}
