package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func OK(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func Created(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusCreated, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func OKMessage(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: message,
	})
}
