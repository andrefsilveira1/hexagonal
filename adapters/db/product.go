package db

import (
	"database/sql"

	"github.com/andrefsilveira1/hexagonal/application"
)

type ProductDB struct {
	db *sql.DB
}

func (p *ProductDB) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("SELECT id, name, price, status from products where id = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
