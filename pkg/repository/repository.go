package repository

import (
	"github.com/jmoiron/sqlx"
	"storeApi"
)

type Store interface {
	CreateProduct(product storeApi.Product) (bool, error)
	GetProducts() ([]storeApi.Product, error)
	GetProductById(productId int) (storeApi.Product, error)
	DeleteProductById(productId int) (storeApi.Product, error)
	UpdateProductById(productId int, product storeApi.Product) (bool, error)
}

type Repository struct {
	Store
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Store: NewStoreMySql(db),
	}
}
