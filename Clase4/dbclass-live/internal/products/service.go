package products

import (
	"dbtest/internal/domain"
)

// constr
func NewService(rp Repository) Service {
	return &service{rp: rp}
}

// controller
type service struct {
	rp Repository
}
// read
func (sv *service) Get() ([]domain.Product, error) {
	return sv.rp.Get()
}
func (sv *service) GetFull(id int) (domain.ProductWithWarehouse, error) {
	return sv.rp.GetFull(id)
}
// write
// ...