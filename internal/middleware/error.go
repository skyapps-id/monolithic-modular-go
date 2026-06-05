package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}

		if appErr, ok := err.(*apperror.AppError); ok {
			logger.ErrorCtx(c.Request().Context(), "request error",
				"method", c.Request().Method,
				"path", c.Request().URL.Path,
				"code", appErr.BizCode,
				"message", appErr.Message,
			)

			resp := map[string]interface{}{
				"success": false,
				"code":    appErr.BizCode,
				"message": appErr.Message,
			}
			if appErr.Details != nil {
				resp["details"] = appErr.Details
			}
			return c.JSON(appErr.HTTPCode, resp)
		}

		if he, ok := err.(*echo.HTTPError); ok {
			return c.JSON(he.Code, map[string]interface{}{
				"success": false,
				"code":    "HTTP_ERROR",
				"message": he.Message.(string),
			})
		}

		logger.ErrorCtx(c.Request().Context(), "unhandled error",
			"method", c.Request().Method,
			"path", c.Request().URL.Path,
			"error", err.Error(),
		)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"code":    "INTERNAL_ERROR",
			"message": "internal server error",
		})
	}
}
