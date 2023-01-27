package products

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}
	
	txdb.Register("txdb", "mysql", fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/warehouse_db", os.Getenv("MYSQL_PSWD")))
}

func TestGetOk(t *testing.T) {
	// arrange
	db, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	defer db.Close()
	
	rp := NewRepositorySQL(db)

	// act
	products, err := rp.Get()

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	t.Logf("products: %+v", products)
}

func TestGetFull_Ok(t *testing.T) {
	// arrange
	db, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	defer db.Close()

	rp := NewRepositorySQL(db)

	// act
	product, err := rp.GetFull(1)

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, product)
	t.Logf("product with warehouse: %+v", product)
}

func TestGetFull_Err(t *testing.T) {
	// arrange
	db, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	defer db.Close()

	rp := NewRepositorySQL(db)

	// act
	product, err := rp.GetFull(99)

	// assert
	assert.Error(t, err)
	// assert.EqualError(t, err, ErrNotFound.Error())
	assert.ErrorIs(t, err, ErrNotFound)
	assert.Empty(t, product)
	t.Logf("error: %s", err)
}