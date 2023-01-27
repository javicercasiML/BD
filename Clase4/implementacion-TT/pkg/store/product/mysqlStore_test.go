package store

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := godotenv.Load("/Users/jcercasi/proyectos/BD/Clase4/implementacion-TT/cmd/server/.env"); err != nil {
		panic(err)
	}

	txdb.Register("txdb", "mysql", fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/my_db", os.Getenv("MYSQL_PSWD")))
}

func TestReadOk(t *testing.T) {
	// arrange
	db, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	expectedProduct := domain.Product{Id: 1, Name: "Corn Shoots", Quantity: 244, CodeValue: "0009-1111", IsPublished: false, Expiration: "2022-01-08", Price: 23.27, Id_warehouse: 1}

	defer db.Close()

	rp := NewSqlStore(db)

	// act
	products, err := rp.Read(1)

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	assert.Equal(t, expectedProduct, products)
	t.Logf("products: %+v", products)
}

func TestRead_Err(t *testing.T) {
	// arrange
	db, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	defer db.Close()

	rp := NewSqlStore(db)

	// act
	product, err := rp.Read(999)

	// assert
	assert.Error(t, err)
	assert.EqualError(t, err, ErrNotFound.Error())
	//assert.ErrorIs(t, err, ErrNotFound.Erro)
	assert.Empty(t, product)
	t.Logf("error: %s", err)
}

func TestGetAllOk(t *testing.T) {
	// arrange
	db, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	defer db.Close()

	rp := NewSqlStore(db)

	// act
	products, err := rp.GetAll()

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	t.Logf("products: %+v", products)
}
