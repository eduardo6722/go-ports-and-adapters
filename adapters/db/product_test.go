package db_test

import (
	"database/sql"
	"testing"

	"github.com/eduardo6722/go-hexagonal/adapters/db"
	"github.com/eduardo6722/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	productTable := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`
	stmt, err := db.Prepare(productTable)
	if err != nil {
		panic(err)
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	stmt, _ := db.Prepare("INSERT INTO products(id, name, price, status) values(?, ?, ?, ?)")
	stmt.Exec("abc", "Product Test", 10, "disabled")
}

func TestProductDBGet(t *testing.T) {
	setup()
	defer Db.Close()

	productDB := db.NewProductDB(Db)

	product, err := productDB.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "abc", product.GetID())
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, float64(10), product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDBSave(t *testing.T) {
	setup()
	defer Db.Close()

	productDB := db.NewProductDB(Db)

	product := application.NewProduct()
	product.Name = "Product Test 2"
	product.Price = 20

	productResult, err := productDB.Save(product)

	require.Nil(t, err)
	require.Equal(t, "Product Test 2", productResult.GetName())
}

func TestProductDBUpdate(t *testing.T) {
	setup()
	defer Db.Close()

	productDB := db.NewProductDB(Db)

	product := application.NewProduct()
	product.Name = "Product Test 3"
	product.Price = 30

	productResult, err := productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, "Product Test 3", productResult.GetName())

	product.Name = "Product Test 3 Updated"
	product.Price = 40
	productResult, err = productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, float64(40), productResult.GetPrice())
	require.Equal(t, "Product Test 3 Updated", productResult.GetName())
}
