package product_group

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/handler/model"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/response"
)

// Create godoc
//
//	@Summary		Create product group
//	@Description	Create new product group
//	@Tags			product-groups
//	@Accept			json
//	@Produce		json
//	@Param			body	body		object	true	"Product group data"
//	@Success		201		{object}	response.SuccessResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Router			/api/v1/product-groups [post]
func (h *ProductGroupHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var req model.CreateProductGroupRequest
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

	return response.Created(c, model.GroupToResponse(created))
}
