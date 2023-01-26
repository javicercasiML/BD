package products

import (
	"database/sql"
	"dbtest/internal/domain"
	"fmt"
)

const (
	GET_ALL = "SELECT id, name, type, count, price, warehouse_id FROM products;"
	GET_FULL= "SELECT p.id, p.name, p.type, p.count, p.price, p.warehouse_id, w.name, w.address FROM products as p " +
			  "INNER JOIN warehouses as w ON p.warehouse_id = w.id " + 
			  "WHERE p.id = ?;"
)

var (
	ErrSTMT = fmt.Errorf("%w. %s", ErrInternal, "statement has failed")
)

// constr
func NewRepositorySQL(db *sql.DB) Repository {
	return &repositorySQL{db: db}
}


// controller
type repositorySQL struct {
	db *sql.DB
}
func (rp *repositorySQL) Get() ([]domain.Product, error) {
	stmt, err := rp.db.Prepare(GET_ALL)
	if err != nil {
		return nil, ErrSTMT
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, ErrInternal
	}

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(
			&product.ID, &product.Name, &product.Type, &product.Count, &product.Price,
			&product.WarehouseID,
		)
		if err != nil {
			return nil, ErrInternal
		}

		products = append(products, product)
	}

	return products, nil
}
func (rp *repositorySQL) GetFull(id int) (domain.ProductWithWarehouse, error) {
	stmt, err := rp.db.Prepare(GET_FULL)
	if err != nil {
		return domain.ProductWithWarehouse{}, ErrSTMT
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var product domain.ProductWithWarehouse
	if err := row.Scan(
		&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.WarehouseID,
		&product.WarehouseName, &product.WarehouseAddress,
	); err != nil {
		switch err {
		case sql.ErrNoRows:
			return domain.ProductWithWarehouse{}, ErrNotFound
		default:
			return domain.ProductWithWarehouse{}, ErrInternal
		}
	}

	return product, nil
}