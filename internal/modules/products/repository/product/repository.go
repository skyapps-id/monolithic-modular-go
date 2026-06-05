package product

import (
	"database/sql"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product"
)

var _ product.ProductRepository = (*productRepository)(nil)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) product.ProductRepository {
	return &productRepository{db: db}
}
