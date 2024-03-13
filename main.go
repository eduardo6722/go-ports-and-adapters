package main

import (
	"database/sql"

	"github.com/eduardo6722/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"

	db2 "github.com/eduardo6722/go-hexagonal/adapters/db"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	productDbAdapter := db2.NewProductDB(db)
	productService := application.NewProductService(productDbAdapter)

	product, err := productService.Create("Product Test", 10)

	productService.Enable(product)

	if err != nil {
		panic(err.Error())
	}

	print("Product created\n")

}
