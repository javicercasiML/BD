package warehouse

import (
	"errors"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	store "github.com/bootcamp-go/consignas-go-db.git/pkg/store/warehouse"
)

type Repository interface {
	GetAll() ([]domain.Warehouse, error)
	// GetByID busca un producto por su id
	GetByID(id int) (domain.Warehouse, error)
	// Create agrega un nuevo producto
	Create(w domain.Warehouse) (domain.Warehouse, error)
	// Update actualiza un producto
	Update(id int, w domain.Warehouse) (domain.Warehouse, error)
	// Delete elimina un producto
	Delete(id int) error
	GetReport(id int) ([]domain.WarehouseReport, error)
}

type repository struct {
	storage store.StoreWarehouseInterface
}

// NewRepository crea un nuevo repositorio
func NewRepositoryW(storage store.StoreWarehouseInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Warehouse, error) {
	warehouse, err := r.storage.Read(id)
	if err != nil {
		return domain.Warehouse{}, errors.New("warehouse not found")
	}
	return warehouse, nil

}

func (r *repository) GetAll() ([]domain.Warehouse, error) {
	warehouses, err := r.storage.GetAll()
	if err != nil {
		return []domain.Warehouse{}, err
	}
	return warehouses, nil
}

func (r *repository) GetReport(id int) ([]domain.WarehouseReport, error) {
	report, err := r.storage.GetProductReport(id)
	if err != nil {
		return []domain.WarehouseReport{}, err
	}
	return report, nil
}

func (r *repository) Create(w domain.Warehouse) (domain.Warehouse, error) {
	id, err := r.storage.Create(w)
	if err != nil {
		return domain.Warehouse{}, errors.New("error creating warehouse")
	}
	w.Id = id
	return w, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, w domain.Warehouse) (domain.Warehouse, error) {

	err := r.storage.Update(w)
	if err != nil {
		return domain.Warehouse{}, errors.New("error updating warehouse")
	}
	return w, nil
}
