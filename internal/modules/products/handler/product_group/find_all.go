package product_group

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/handler/model"
	"github.com/skyapps-id/monolithic-modular-go/pkg/response"
)

// FindAll godoc
//
//	@Summary		List product groups
//	@Description	Get all product groups
//	@Tags			product-groups
//	@Produce		json
//	@Success		200		{object}	response.SuccessResponse
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/api/v1/product-groups [get]
func (h *ProductGroupHandler) FindAll(c echo.Context) error {
	ctx := c.Request().Context()

	groups, err := h.uc.FindAll(ctx)
	if err != nil {
		return err
	}

	return response.OK(c, model.GroupsToResponse(groups))
}
