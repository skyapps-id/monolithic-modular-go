package users

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	domainuser "github.com/skyapps-id/monolithic-modular-go/internal/modules/users/domain/user"
	httpuser "github.com/skyapps-id/monolithic-modular-go/internal/modules/users/handler/user"
	repouser "github.com/skyapps-id/monolithic-modular-go/internal/modules/users/repository/user"
	ucuser "github.com/skyapps-id/monolithic-modular-go/internal/modules/users/usecase/user"
	"github.com/skyapps-id/monolithic-modular-go/internal/router"
)

var _ router.Module = (*Module)(nil)

type Module struct {
	Handler *httpuser.UserHandler
	Repo    domainuser.UserRepository
}

func NewModule(db *sql.DB) *Module {
	repo := repouser.NewUserRepository(db)
	uc := ucuser.NewUserUsecase(repo)
	h := httpuser.NewUserHandler(uc)
	return &Module{Handler: h, Repo: repo}
}

func (m *Module) Name() string {
	return "users"
}

func (m *Module) RegisterRoutes(g *echo.Group) {
	users := g.Group("/users")
	users.POST("", m.Handler.Create)
	users.GET("", m.Handler.FindAll)
	users.GET("/:id", m.Handler.FindByID)
}
