package product

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/handler/model"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/response"
)

// Create godoc
//
//	@Summary		Create product
//	@Description	Create new product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			body	body		object	true	"Product data"
//	@Success		201		{object}	response.SuccessResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Router			/api/v1/products [post]
func (h *ProductHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var req model.CreateProductRequest
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

	return response.Created(c, model.ProductToResponse(created))
}
