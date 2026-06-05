package product_group

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/handler/model"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/response"
)

// FindByID godoc
//
//	@Summary		Get product group
//	@Description	Get product group by ID
//	@Tags			product-groups
//	@Produce		json
//	@Param			id	path		string	true	"Product group ID"
//	@Success		200		{object}	response.SuccessResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		404		{object}	response.ErrorResponse
//	@Router			/api/v1/product-groups/{id} [get]
func (h *ProductGroupHandler) FindByID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	if id == "" {
		return apperror.ValidationError("id is required")
	}

	found, err := h.uc.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return response.OK(c, model.GroupToResponse(found))
}
