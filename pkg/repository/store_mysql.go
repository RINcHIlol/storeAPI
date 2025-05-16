package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"storeApi"
)

type StoreMySql struct {
	db *sqlx.DB
}

func NewStoreMySql(db *sqlx.DB) *StoreMySql {
	return &StoreMySql{db: db}
}

func (p *StoreMySql) CreateProduct(product storeApi.Product) (bool, error) {
	query := `
		INSERT INTO products (order_num, name, price, description, image)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := p.db.Exec(query, product.OrderNum, product.Name, product.Price, product.Description, product.Image)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *StoreMySql) GetProducts() ([]storeApi.Product, error) {
	var products []storeApi.Product

	query := `SELECT id, order_num, name, price, description, image FROM products`

	err := p.db.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *StoreMySql) GetProductById(productId int) (storeApi.Product, error) {
	var product storeApi.Product

	query := `SELECT id, order_num, name, price, description, image FROM products WHERE id = ?`

	err := p.db.Get(&product, query, productId)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *StoreMySql) DeleteProductById(productId int) (storeApi.Product, error) {
	var product storeApi.Product

	querySelect := `SELECT id, order_num, name, price, description, image FROM products WHERE id = ?`
	err := p.db.Get(&product, querySelect, productId)
	if err != nil {
		return storeApi.Product{}, err
	}

	queryDelete := `DELETE FROM products WHERE id = ?`
	result, err := p.db.Exec(queryDelete, productId)
	if err != nil {
		return storeApi.Product{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return storeApi.Product{}, err
	}
	if rowsAffected == 0 {
		return storeApi.Product{}, fmt.Errorf("no product with id %d found", productId)
	}

	return product, nil
}

func (p *StoreMySql) UpdateProductById(productId int, product storeApi.Product) (bool, error) {
	query := `
		UPDATE products 
		SET order_num = ?, name = ?, price = ?, description = ?, image = ?
		WHERE id = ?
	`

	result, err := p.db.Exec(query, product.OrderNum, product.Name, product.Price, product.Description, product.Image, productId)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
