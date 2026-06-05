package middleware

import (
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

func TraceID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		traceID := c.Request().Header.Get("X-Trace-ID")
		if traceID == "" {
			traceID = uuid.New().String()
			traceID = strings.ReplaceAll(traceID, "-", "")
		}
		ctx := logger.WithTraceID(c.Request().Context(), traceID)
		c.SetRequest(c.Request().WithContext(ctx))
		c.Response().Header().Set("X-Trace-ID", traceID)
		return next(c)
	}
}
