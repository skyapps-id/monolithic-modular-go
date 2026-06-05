package product

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/handler/model"
	"github.com/skyapps-id/monolithic-modular-go/pkg/response"
)

// FindAll godoc
//
//	@Summary		List products
//	@Description	Get all products
//	@Tags			products
//	@Produce		json
//	@Success		200		{object}	response.SuccessResponse
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/api/v1/products [get]
func (h *ProductHandler) FindAll(c echo.Context) error {
	ctx := c.Request().Context()

	products, err := h.uc.FindAll(ctx)
	if err != nil {
		return err
	}

	return response.OK(c, model.ProductsToResponse(products))
}
