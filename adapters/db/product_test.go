package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/andrefsilveira1/hexagonal/adapters/db"
	"github.com/andrefsilveira1/hexagonal/application"
	"github.com/stretchr/testify/require"
)

var DB *sql.DB

func setup() {
	DB, _ = sql.Open("sqlite3", ":memory:")

	createTable(DB)
	createProduct(DB)
}

func createTable(database *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`

	stmt, err := database.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(database *sql.DB) {
	insert := `INSERT INTO products values("abc", "Product Test", 0, "disabled")`
	stmt, err := database.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDB(t *testing.T) {
	setup()
	defer DB.Close()
	productDB := db.NewProductDB(DB)
	product, err := productDB.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
}

func TestProductDB_Save(t *testing.T) {
	setup()
	defer DB.Close()
	productDb := db.NewProductDB(DB)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 30

	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}
