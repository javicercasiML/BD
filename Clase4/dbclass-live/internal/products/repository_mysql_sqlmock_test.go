package products

import (
	"database/sql"
	"dbtest/internal/domain"
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

	data := []domain.Product{
		{ID: 1, Name: "Coca-Poca", Type: "Bebidas", Count: 5, Price: 105.5, WarehouseID: 1},
		{ID: 2, Name: "Freezo-Ice", Type: "Helados", Count: 7, Price: 75, WarehouseID: 1},
		{ID: 3, Name: "Papas-PLays", Type: "Snacks", Count: 3, Price: 40, WarehouseID: 2},
	}
	rows := mock.NewRows([]string{"id", "name", "type", "count", "price", "warehouse_id"})
	for _, d := range data {
		rows.AddRow(d.ID, d.Name, d.Type, d.Count, d.Price, d.WarehouseID)
	}

	mock.ExpectPrepare(regexp.QuoteMeta(GET_ALL)).ExpectQuery().WillReturnRows(rows)

	rp := NewRepositorySQL(db)

	// act
	products, err := rp.Get()

	// assert
	assert.NoError(t, err)
	assert.Equal(t, data, products)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGet_SQLMock_Err(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// data := []domain.Product{
	// 	{ID: 1, Name: "Coca-Poca", Type: "Bebidas", Count: 5, Price: 105.5, WarehouseID: 1},
	// 	{ID: 2, Name: "Freezo-Ice", Type: "Helados", Count: 7, Price: 75, WarehouseID: 1},
	// 	{ID: 3, Name: "Papas-PLays", Type: "Snacks", Count: 3, Price: 40, WarehouseID: 2},
	// }
	// rows := mock.NewRows([]string{"id", "name", "type", "count", "price", "warehouse_id"})
	// for _, d := range data {
	// 	rows.AddRow(d.ID, d.Name, d.Type, d.Count, d.Price, d.WarehouseID)
	// }

	mock.ExpectPrepare(regexp.QuoteMeta(GET_ALL)).ExpectQuery().WillReturnError(sql.ErrConnDone)

	rp := NewRepositorySQL(db)

	// act
	products, err := rp.Get()

	// assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInternal)
	assert.Empty(t, products)
	assert.NoError(t, mock.ExpectationsWereMet())
}