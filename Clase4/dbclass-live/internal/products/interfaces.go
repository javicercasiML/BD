package products

import (
	"dbtest/internal/domain"
	"errors"
)

// --------------------------------------------------------------------------------------------------------
// Interface Service
type Service interface {
	// read
	Get() ([]domain.Product, error)
	GetFull(id int) (domain.ProductWithWarehouse, error)
	// write
	// ...
}

var (
	ErrValidation = errors.New("movie validation failed")
)


// --------------------------------------------------------------------------------------------------------
// Interface Repository
type Repository interface {
	// read
	Get() ([]domain.Product, error)
	GetFull(id int) (domain.ProductWithWarehouse, error)
	// write
	// ...
}

var (
	ErrInternal = errors.New("there was an error in the internal implementation")
	ErrNotFound = errors.New("resource was not found")
)