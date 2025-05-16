package service

import (
	"storeApi"
	"storeApi/pkg/repository"
)

type Store interface {
	AddNewProduct(product storeApi.Product) (bool, error)
	GetProducts() ([]storeApi.Product, error)
	GetProductById(productId int) (storeApi.Product, error)
	BuyProduct(productId int) (bool, error)
	UpdateProductById(productId int, product storeApi.Product) (bool, error)
}

type Service struct {
	Store
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Store: NewStoreService(repos.Store),
	}
}
