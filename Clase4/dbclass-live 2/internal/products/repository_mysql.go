package products

import (
	"database/sql"
	"dbtest/internal/domain"
	"fmt"
)

const (
	// read
	GET_ALL = "SELECT id, name, type, count, price, warehouse_id FROM products;"
	GET_FULL= "SELECT p.id, p.name, p.type, p.count, p.price, p.warehouse_id, w.name, w.address FROM products as p " +
			  "INNER JOIN warehouses as w ON p.warehouse_id = w.id " + 
			  "WHERE p.id = ?;"
	// write
	CREATE  = "INSERT INTO products (name, type, count, price, warehouse_id) VALUES (?, ?, ?, ?, ?);"
)

var (
	ErrSTMT = fmt.Errorf("%w. %s", ErrInternal, "statement has failed")
	ErrExec = fmt.Errorf("%w. %s", ErrInternal, "execution has failed")
	ErrRowsAffected = fmt.Errorf("%w. %s", ErrInternal, "more than 1 row affected")
)

// constr
func NewRepositorySQL(db *sql.DB) Repository {
	return &repositorySQL{db: db}
}


// controller
type repositorySQL struct {
	db *sql.DB
}
// read
func (rp *repositorySQL) Get() ([]domain.Product, error) {
	stmt, err := rp.db.Prepare(GET_ALL)
	if err != nil {
		return nil, fmt.Errorf("%w. %s", ErrSTMT, err)
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
// write
func (rp *repositorySQL) Create(product domain.Product) (int, error) {
	stmt, err := rp.db.Prepare(CREATE)
	if err != nil {
		return 0, ErrSTMT
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Name, product.Type, product.Count, product.Price, product.WarehouseID)
	if err != nil {
		return 0, ErrExec
	}

	// check results
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, ErrInternal
	}

	if rows != 1 {
		return 0, ErrRowsAffected
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, ErrInternal
	}

	return int(lastId), nil
}