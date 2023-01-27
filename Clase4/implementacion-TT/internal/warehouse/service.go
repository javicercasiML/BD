package warehouse

import (
	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Warehouse, error)
	// GetByID busca un producto por su id
	GetByID(id int) (domain.Warehouse, error)
	// Create agrega un nuevo producto
	Create(w domain.Warehouse) (domain.Warehouse, error)
	// Delete elimina un producto
	Delete(id int) error
	// Update actualiza un producto
	Update(id int, u domain.Warehouse) (domain.Warehouse, error)

	GetReport(id int) ([]domain.WarehouseReport, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewServiceW(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Warehouse, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Warehouse{}, err
	}
	return p, nil
}

func (s *service) GetAll() ([]domain.Warehouse, error) {
	w, err := s.r.GetAll()
	if err != nil {
		return []domain.Warehouse{}, err
	}
	return w, nil
}

func (s *service) GetReport(id int) ([]domain.WarehouseReport, error) {
	report, err := s.r.GetReport(id)
	if err != nil {
		return []domain.WarehouseReport{}, err
	}
	return report, nil
}

func (s *service) Create(w domain.Warehouse) (domain.Warehouse, error) {
	p, err := s.r.Create(w)
	if err != nil {
		return domain.Warehouse{}, err
	}
	return p, nil
}
func (s *service) Update(id int, u domain.Warehouse) (domain.Warehouse, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Warehouse{}, err
	}
	if u.Name != "" {
		p.Name = u.Name
	}
	if u.Adress != "" {
		p.Adress = u.Adress
	}
	if u.Telephone != "" {
		p.Telephone = u.Telephone
	}
	if u.Capacity > 0 {
		p.Capacity = u.Capacity
	}

	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Warehouse{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
