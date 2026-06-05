package product

type ProductRepository interface {
	Save(product *Product) error
	FindByID(id string) (*Product, error)
	FindAll() ([]Product, error)
}
