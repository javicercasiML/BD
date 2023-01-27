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
	expectedWarehouse := domain.Warehouse{Id: 1, Name: "Main Warehouse", Adress: "221 Baker Street", Telephone: "4555666", Capacity: 100}
	defer db.Close()

	rp := NewSqlStoreW(db)

	// act
	warehouses, err := rp.Read(1)

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, warehouses)
	assert.Equal(t, expectedWarehouse, warehouses)
	t.Logf("warehouses: %+v", warehouses)
}

func TestGetAllOk(t *testing.T) {
	// arrange
	db, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	defer db.Close()

	rp := NewSqlStoreW(db)

	// act
	warehouses, err := rp.GetAll()

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, warehouses)
	t.Logf("warehouses: %+v", warehouses)
}

func TestCreateOk(t *testing.T) {
	// arrange
	db, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	sendWarehouse := domain.Warehouse{Name: "Main Warehouse", Adress: "221 Baker Street", Telephone: "4555666", Capacity: 100}

	defer db.Close()

	rp := NewSqlStoreW(db)

	// act
	warehouse, err := rp.Create(sendWarehouse)

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, warehouse)
	t.Logf("ID warehouse: %+v", warehouse)
}
