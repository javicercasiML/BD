package main

import (
	"database/sql"
	"dbtest/internal/products"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// ------------------------
	// env
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	
	// ------------------------
	// instances
	// db := []domain.Product{
	// 	{ID: 1, Name: "Coca-Poca", Type: "Bebidas", Count: 5, Price: 105.5, WarehouseID: 1},
	// 	{ID: 2, Name: "Freezo-Ice", Type: "Helados", Count: 7, Price: 75, WarehouseID: 1},
	// 	{ID: 3, Name: "Papas-PLays", Type: "Snacks", Count: 3, Price: 40, WarehouseID: 2},
	// }
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/warehouse_db", os.Getenv("MYSQL_PSWD")))
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}


	// rp := products.NewRepositoryLocal(&db, 3)
	rp := products.NewRepositorySQL(db)
	sv := products.NewService(rp)

	// ------------------------
	// app
	// -> get
	// mvs, err := sv.Get()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("-products: %+v\n", mvs)

	mvs, err := sv.GetFull(2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("-products with warehouse: %+v\n", mvs)
}