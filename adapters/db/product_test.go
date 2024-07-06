package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/andrefsilveira1/hexagonal/adapters/db"
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
