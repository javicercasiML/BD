package store

import (
	"database/sql"
	"fmt"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/go-sql-driver/mysql"
)

type mysqlStore struct {
	db *sql.DB
}

func NewSqlStoreW(db *sql.DB) StoreWarehouseInterface {
	return &mysqlStore{db: db}
}

const (
	GET_ALL   = "SELECT id, name, adress, telephone, capacity FROM warehouses;"
	GetReport = "SELECT w.name, COUNT(p.id_warehouse) as product_count FROM warehouses as w " +
		"LEFT JOIN products as p ON w.id = p.id_warehouse " +
		"GROUP BY w.id;"
	GetReportQuery = "SELECT w.name, COUNT(p.id_warehouse) as product_count FROM warehouses as w " +
		"LEFT JOIN products as p ON w.id = p.id_warehouse " +
		"GROUP BY w.id HAVING w.id = ?;"
)

var (
	ErrSTMT = fmt.Errorf("%w. %s", ErrInternal, "statement has failed")
)

// Get searches for a warehouse by its id
func (m *mysqlStore) Read(id int) (domain.Warehouse, error) {
	query := "SELECT id, name, adress, telephone, capacity FROM warehouses WHERE id=?;"
	row := m.db.QueryRow(query, id)

	if row.Err() != nil {
		switch row.Err() {
		case sql.ErrNoRows:
			return domain.Warehouse{}, ErrNotFound
		default:
			return domain.Warehouse{}, ErrInternal
		}
	}

	w := domain.Warehouse{}
	err := row.Scan(&w.Id, &w.Name, &w.Adress, &w.Telephone, &w.Capacity)
	if err != nil {
		return domain.Warehouse{}, ErrInternal
	}

	return w, nil
}

func (m *mysqlStore) GetAll() ([]domain.Warehouse, error) {
	stmt, err := m.db.Prepare(GET_ALL)
	if err != nil {
		return nil, ErrSTMT
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, ErrInternal
	}

	var warehouses []domain.Warehouse
	for rows.Next() {
		var w domain.Warehouse
		err := rows.Scan(&w.Id, &w.Name, &w.Adress, &w.Telephone, &w.Capacity)

		if err != nil {
			return nil, ErrInternal
		}

		warehouses = append(warehouses, w)
	}

	return warehouses, nil
}

func (m *mysqlStore) GetProductReport(id int) ([]domain.WarehouseReport, error) {
	var rows *sql.Rows
	var err error
	if id > 0 {
		rows, err = m.db.Query(GetReportQuery, id)
	} else {
		rows, err = m.db.Query(GetReport)
	}
	if err != nil {
		return nil, ErrInternal
	}

	var reports []domain.WarehouseReport
	for rows.Next() {
		var report domain.WarehouseReport
		err := rows.Scan(&report.Name, &report.ProductCount)
		if err != nil {
			return nil, ErrInternal
		}
		reports = append(reports, report)
	}
	return reports, nil
}

// Save adds a new warehouse to the database
func (m *mysqlStore) Create(warehouse domain.Warehouse) (int, error) {
	query := "INSERT INTO warehouses (name, adress, telephone, capacity) VALUES (?, ?, ?, ?);"
	stmt, err := m.db.Prepare(query)

	if err != nil {
		return 0, ErrInternal
	}

	defer stmt.Close()

	res, err := stmt.Exec(&warehouse.Name, &warehouse.Adress, &warehouse.Telephone, &warehouse.Capacity)
	if err != nil {
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			return 0, ErrInternal
		}
		switch driverErr.Number {
		case 1062:
			return 0, ErrDuplicated
		default:
			return 0, ErrInternal
		}
	}

	RowsAffected, err := res.RowsAffected()
	if err != nil || RowsAffected != 1 {
		return 0, ErrInternal
	}

	warehouseID, err := res.LastInsertId()
	if err != nil {
		return 0, ErrInternal
	}

	warehouse.Id = int(warehouseID)

	return int(warehouseID), nil
}

// Update updates a warehouse from the database.
func (m *mysqlStore) Update(warehouse domain.Warehouse) error {
	query := "UPDATE products SET name = ?, adress = ?, telephone = ?, capacity = ? WHERE id = ?;"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return ErrInternal
	}

	res, err := stmt.Exec(&warehouse.Name, &warehouse.Adress, &warehouse.Telephone, &warehouse.Capacity, &warehouse.Id)
	if err != nil {
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			return ErrInternal

		}
		switch driverErr.Number {
		case 1072:
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
	query := "DELETE FROM warehouses WHERE id=?"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return ErrInternal
	}

	res, err := stmt.Exec(id)
	println("Error: ", err)
	if err != nil {
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			return ErrInternal
		}
		switch driverErr.Number {
		case 1072:
			return ErrNotFound
		default:
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
