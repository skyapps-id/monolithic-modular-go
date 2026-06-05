package user

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/users/handler/model"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/response"
)

// FindByID godoc
//
//	@Summary		Get user
//	@Description	Get user by ID
//	@Tags			users
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200		{object}	response.SuccessResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		404		{object}	response.ErrorResponse
//	@Router			/api/v1/users/{id} [get]
func (h *UserHandler) FindByID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	if id == "" {
		return apperror.ValidationError("id is required")
	}

	found, err := h.uc.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return response.OK(c, model.UserToResponse(found))
}
