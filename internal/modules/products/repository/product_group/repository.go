package product_group

import (
	"database/sql"

	"github.com/skyapps-id/monolithic-modular-go/internal/modules/products/domain/product_group"
)

var _ product_group.ProductGroupRepository = (*productGroupRepository)(nil)

type productGroupRepository struct {
	db *sql.DB
}

func NewProductGroupRepository(db *sql.DB) product_group.ProductGroupRepository {
	return &productGroupRepository{db: db}
}
