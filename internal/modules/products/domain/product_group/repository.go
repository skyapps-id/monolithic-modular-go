package product_group

type ProductGroupRepository interface {
	Save(group *ProductGroup) error
	FindByID(id string) (*ProductGroup, error)
	FindAll() ([]ProductGroup, error)
}
