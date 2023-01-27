package store

import (
	"errors"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

var (
	ErrNotFound   = errors.New("warehouse not found")
	ErrInternal   = errors.New("an internal error")
	ErrDuplicated = errors.New("warehouse already exists")
)

type StoreWarehouseInterface interface {
	GetAll() ([]domain.Warehouse, error)
	// Read devuelve un producto por su id
	Read(id int) (domain.Warehouse, error)
	// Create agrega un nuevo producto
	GetProductReport(id int) ([]domain.WarehouseReport, error)

	Create(warehouse domain.Warehouse) (int, error)
	// Update actualiza un producto
	Update(warehouse domain.Warehouse) error
	// Delete elimina un producto
	Delete(id int) error
	// Exists verifica si un producto existe
}
