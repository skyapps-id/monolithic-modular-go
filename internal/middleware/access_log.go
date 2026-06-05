package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

func AccessLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		path := c.Request().URL.Path
		method := c.Request().Method

		err := next(c)

		duration := time.Since(start)
		status := c.Response().Status
		if err != nil {
			status = 500
		}

		logger.InfoCtx(c.Request().Context(), "http request",
			"method", method,
			"path", path,
			"status", status,
			"duration", duration.String(),
		)

		return err
	}
}
