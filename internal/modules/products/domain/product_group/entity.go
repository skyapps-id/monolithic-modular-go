package product_group

import "time"

type ProductGroup struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
