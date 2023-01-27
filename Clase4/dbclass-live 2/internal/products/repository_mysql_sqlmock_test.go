package products

import (
	"database/sql"
	"dbtest/internal/domain"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGet_SQLMock_Ok(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	expected := []domain.Product{
		{ID: 1, Name: "Coca-Poca", Type: "Bebidas", Count: 5, Price: 105.5, WarehouseID: 1},
		{ID: 2, Name: "Freezo-Ice", Type: "Helados", Count: 7, Price: 75, WarehouseID: 1},
		{ID: 3, Name: "Papas-PLays", Type: "Snacks", Count: 3, Price: 40, WarehouseID: 2},
	}
	rows := mock.NewRows([]string{"id", "name", "type", "count", "price", "warehouse_id"})
	for _, d := range expected {
		rows.AddRow(d.ID, d.Name, d.Type, d.Count, d.Price, d.WarehouseID)
	}

	mock.ExpectPrepare(
		regexp.QuoteMeta("SELECT id, name, type, count, price, warehouse_id FROM products;"),
		).
		ExpectQuery().WillReturnRows(rows)
	

	rp := NewRepositorySQL(db)

	// act
	products, err := rp.Get()

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, products)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGet_SQLMock_ErrPrepare(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(GET_ALL)).WillReturnError(sql.ErrConnDone)
	
	rp := NewRepositorySQL(db)

	// act
	products, err := rp.Get()

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrSTMT)
	// assert.EqualError(t, err, ErrSTMT.Error())
	assert.Empty(t, products)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGet_SQLMock_ErrQuery(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(GET_ALL)).
	ExpectQuery().WillReturnError(errors.New("sql error"))
	
	rp := NewRepositorySQL(db)

	// act
	products, err := rp.Get()

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInternal)
	// assert.EqualError(t, err, ErrSTMT.Error())
	assert.Empty(t, products)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// write
func TestCreate_SQLMock_Ok(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(CREATE)).
	ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))

	rp := NewRepositorySQL(db)
	
	dp := domain.Product{
		Name:        "Pepsi",
		Type:        "Freeze",
		Count:       150,
		Price:       190,
		WarehouseID: 1,
	}

	// act
	lastId, err := rp.Create(dp)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 1, lastId)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreate_SQLMock_Err(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dp := domain.Product{
		Name:        "Pepsi",
		Type:        "Freeze",
		Count:       150,
		Price:       190,
		WarehouseID: 1,
	}

	mock.ExpectPrepare(regexp.QuoteMeta(CREATE)).
	ExpectExec().WithArgs(dp.Name, dp.Type, dp.Count, dp.Price, dp.WarehouseID).
	WillReturnError(sql.ErrConnDone)

	rp := NewRepositorySQL(db)

	// act
	lastId, err := rp.Create(dp)

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInternal)
	assert.Equal(t, 0, lastId)
	assert.NoError(t, mock.ExpectationsWereMet())
}