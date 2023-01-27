package products

import "dbtest/internal/domain"

// constr
func NewRepositoryLocal(db *[]domain.Product, lastId int) Repository {
	return &repositoryLocal{db: db, lastId: lastId}
}

// controller
type repositoryLocal struct {
	db 	   *[]domain.Product
	lastId int
}
// read
func (rp *repositoryLocal) Get() ([]domain.Product, error) {
	return *rp.db, nil
}
func (rp *repositoryLocal) GetFull(id int) (domain.ProductWithWarehouse, error) {
	return domain.ProductWithWarehouse{}, nil
}
// write
func (rp *repositoryLocal) Create(product domain.Product) (int, error) {
	return 0, nil
}