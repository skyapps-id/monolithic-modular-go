package product

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/handler/model"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/response"
)

// FindByID godoc
//
//	@Summary		Get product
//	@Description	Get product by ID
//	@Tags			products
//	@Produce		json
//	@Param			id	path		string	true	"Product ID"
//	@Success		200		{object}	response.SuccessResponse
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		404		{object}	response.ErrorResponse
//	@Router			/api/v1/products/{id} [get]
func (h *ProductHandler) FindByID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	if id == "" {
		return apperror.ValidationError("id is required")
	}

	found, err := h.uc.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return response.OK(c, model.ProductToResponse(found))
}
