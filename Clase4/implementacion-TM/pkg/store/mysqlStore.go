package store

import (
	"database/sql"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/go-sql-driver/mysql"
)

type mysqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &mysqlStore{db: db}
}

// Get searches for a warehouse by its id
func (m *mysqlStore) Read(id int) (domain.Product, error) {
	query := "SELECT id, name, quantity, code_value, is_published, expiration, price  FROM products WHERE id=?;"
	row := m.db.QueryRow(query, id)

	if row.Err() != nil {
		switch row.Err() {
		case sql.ErrNoRows:
			return domain.Product{}, ErrNotFound
		default:
			return domain.Product{}, ErrInternal
		}
	}

	p := domain.Product{}
	err := row.Scan(&p.Id, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price)
	if err != nil {
		return domain.Product{}, ErrInternal
	}

	return p, nil
}

// Exist validates that the warehouseCode doesn't exist in the database.
func (m *mysqlStore) Exists(codeValue string) bool {
	query := "SELECT id FROM products WHERE code_value = ?;"
	row := m.db.QueryRow(query, codeValue)
	err := row.Scan(&codeValue)
	if err != nil {
		return false
	}
	return true
}

// Save adds a new warehouse to the database
func (m *mysqlStore) Create(product domain.Product) error {
	query := "INSERT INTO products (id, name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?, ?);"
	stmt, err := m.db.Prepare(query)

	if err != nil {
		return ErrInternal
	}

	defer stmt.Close()

	res, err := stmt.Exec(&product.Id, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
	if err != nil {
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			return ErrInternal
		}
		switch driverErr.Number {
		case 1062:
			return ErrDuplicated
		default:
			return ErrInternal
		}
	}

	RowsAffected, err := res.RowsAffected()
	if err != nil || RowsAffected != 1 {
		return ErrInternal
	}

	productID, err := res.LastInsertId()
	if err != nil {
		return ErrInternal
	}

	product.Id = int(productID)

	return nil
}

// Update updates a warehouse from the database.
func (m *mysqlStore) Update(product domain.Product) error {
	query := "UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?;"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return ErrInternal
	}

	res, err := stmt.Exec(&product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price, &product.Id)
	if err != nil {
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			return ErrInternal

		}
		switch driverErr.Number {
		case 1032:
			return ErrNotFound
		default:
			return ErrInternal
		}
	}

	RowsAffected, err := res.RowsAffected()
	if err != nil {
		return ErrInternal
	}

	if RowsAffected != 1 {
		return ErrNotFound
	}
	return nil
}

// Delete removes a product from the database.
func (m *mysqlStore) Delete(id int) error {
	query := "DELETE FROM products WHERE id=?"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return ErrInternal
	}

	res, err := stmt.Exec(id)
	println(err)
	if err != nil {
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			return ErrInternal
		}
		if driverErr.Number != 0 {
			return ErrInternal
		}
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return ErrInternal
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}
