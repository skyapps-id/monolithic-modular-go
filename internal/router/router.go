package router

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/internal/middleware"
	"github.com/skyapps-id/monolithic-modular-go/pkg/validator"
	"github.com/swaggo/echo-swagger"
)

type Module interface {
	Name() string
	RegisterRoutes(g *echo.Group)
}

type App struct {
	echo      *echo.Echo
	modules   []Module
	apiPrefix string
}

func NewApp(e *echo.Echo, apiPrefix string) *App {
	e.Validator = validator.New()

	e.Use(middleware.TraceID)
	e.Use(middleware.AccessLog)
	e.Use(middleware.ErrorHandler)

	e.GET("/swagger/doc.json", func(c echo.Context) error {
		return c.File("docs/swagger.json")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return &App{
		echo:      e,
		apiPrefix: apiPrefix,
	}
}

func (a *App) Register(m ...Module) {
	a.modules = append(a.modules, m...)
}

func (a *App) Modules() []Module {
	return a.modules
}

func (a *App) Serve(addr string) error {
	api := a.echo.Group(a.apiPrefix)
	for _, m := range a.modules {
		m.RegisterRoutes(api)
	}
	return a.echo.Start(addr)
}
