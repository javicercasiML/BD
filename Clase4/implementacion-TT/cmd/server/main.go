package main

import (
	"database/sql"

	"github.com/bootcamp-go/consignas-go-db.git/cmd/server/handler"
	"github.com/bootcamp-go/consignas-go-db.git/internal/product"
	"github.com/bootcamp-go/consignas-go-db.git/internal/warehouse"
	store "github.com/bootcamp-go/consignas-go-db.git/pkg/store/product"
	storeW "github.com/bootcamp-go/consignas-go-db.git/pkg/store/warehouse"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func main() {

	//storage := store.NewJsonStore("./products.json")
	databaseConfig := &mysql.Config{
		User:   "root",
		Passwd: "",
		Addr:   "localhost:3306",
		DBName: "my_db"}

	db, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	// Ping database connection.
	if err = db.Ping(); err != nil {
		panic(err)
	}

	r := gin.Default()

	storageProduct := store.NewSqlStore(db)
	repoProduct := product.NewRepository(storageProduct)
	serviceProduct := product.NewService(repoProduct)
	productHandler := handler.NewProductHandler(serviceProduct)

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET(":id", productHandler.GetByID())
		products.GET("", productHandler.GetAll())
		products.POST("", productHandler.Post())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())
		products.PUT(":id", productHandler.Put())
	}

	storageWarehouse := storeW.NewSqlStoreW(db)
	repoWarehouse := warehouse.NewRepositoryW(storageWarehouse)
	serviceWarehouse := warehouse.NewServiceW(repoWarehouse)
	warehouseHandler := handler.NewWarehouseHandler(serviceWarehouse)
	warehouses := r.Group("/warehouses")
	{
		warehouses.GET(":id", warehouseHandler.GetByID())
		warehouses.GET("", warehouseHandler.GetAll())
		warehouses.POST("", warehouseHandler.Post())
		warehouses.DELETE(":id", warehouseHandler.Delete())
		warehouses.PATCH(":id", warehouseHandler.Patch())
		warehouses.PUT(":id", warehouseHandler.Put())
		warehouses.GET("/reportProducts", warehouseHandler.GetReport())
	}

	r.Run(":8080")
}
