package products

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	httpproduct "github.com/skyapps-id/monolithic-modular-go/internal/modules/products/handler/product"
	httpgroup "github.com/skyapps-id/monolithic-modular-go/internal/modules/products/handler/product_group"
	repoproduct "github.com/skyapps-id/monolithic-modular-go/internal/modules/products/repository/product"
	repogroup "github.com/skyapps-id/monolithic-modular-go/internal/modules/products/repository/product_group"
	ucproduct "github.com/skyapps-id/monolithic-modular-go/internal/modules/products/usecase/product"
	ucgroup "github.com/skyapps-id/monolithic-modular-go/internal/modules/products/usecase/product_group"
	"github.com/skyapps-id/monolithic-modular-go/internal/router"
)

var _ router.Module = (*Module)(nil)

type Module struct {
	ProductHandler *httpproduct.ProductHandler
	GroupHandler   *httpgroup.ProductGroupHandler
}

func NewModule(db *sql.DB) *Module {
	productRepo := repoproduct.NewProductRepository(db)
	productUC := ucproduct.NewProductUsecase(productRepo)
	productH := httpproduct.NewProductHandler(productUC)

	groupRepo := repogroup.NewProductGroupRepository(db)
	groupUC := ucgroup.NewProductGroupUsecase(groupRepo)
	groupH := httpgroup.NewProductGroupHandler(groupUC)

	return &Module{
		ProductHandler: productH,
		GroupHandler:   groupH,
	}
}

func (m *Module) Name() string {
	return "products"
}

func (m *Module) RegisterRoutes(g *echo.Group) {
	products := g.Group("/products")
	products.POST("", m.ProductHandler.Create)
	products.GET("", m.ProductHandler.FindAll)
	products.GET("/:id", m.ProductHandler.FindByID)

	groups := g.Group("/product-groups")
	groups.POST("", m.GroupHandler.Create)
	groups.GET("", m.GroupHandler.FindAll)
	groups.GET("/:id", m.GroupHandler.FindByID)
}
