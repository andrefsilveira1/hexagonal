package main

import (
	"database/sql"
	"log"

	"github.com/andrefsilveira1/hexagonal/adapters/db"
	"github.com/andrefsilveira1/hexagonal/application"
)

func main() {
	Db, _ := sql.Open("sqlite3", "db.sqlite")
	productDb := db.NewProductDB(Db)
	productService := application.NewProductService(productDb)
	product, err := productService.Create("Example", 50)
	if err != nil {
		log.Fatal(err.Error())
	}

	productService.Enable(product)
}
