package db

import (
	"database/sql"

	"github.com/eduardo6722/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db}
}

func (p *ProductDB) Get(id string) (application.IProduct, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDB) Save(product application.IProduct) (application.IProduct, error) {
	var rows int
	err := p.db.QueryRow("select count(id) from products where id = ?", product.GetID()).Scan(&rows)
	if err != nil {
		return nil, err
	}
	if rows > 0 {
		return p.update(product)
	}
	return p.create(product)
}

func (p *ProductDB) create(product application.IProduct) (application.IProduct, error) {
	stmt, err := p.db.Prepare("insert into products(id, name, price, status) values(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDB) update(product application.IProduct) (application.IProduct, error) {
	stmt, err := p.db.Prepare("update products set name = ?, price = ?, status = ? where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}
	return product, nil
}
