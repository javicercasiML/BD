package main

//package db
import (
	"database/sql"
	"log"

	// go get "github.com/go-sql-driver/mysql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

func main() {
	dataSource := "root:@tcp(localhost:3306)/storage"
	// Open inicia un pool de conexiones. Sólo abrir una vez
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database Configured")
}
