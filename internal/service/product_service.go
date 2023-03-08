package service

import (
	"sirawit/shop/internal/config"
	"sirawit/shop/internal/model"
	"sirawit/shop/internal/repository"
)

type ProductService interface {
	GetProducts(id uint64) ([]model.Product, error)
	CreateProduct(input model.Product) (*model.Product, error)
	GetProduct(id uint64) (*model.Product, error)
}

type productService struct {
	db     repository.ProductQuery
	config config.ProductConfig
}

func NewProductService(db repository.ProductQuery, config config.ProductConfig) ProductService {
	return &productService{
		db:     db,
		config: config,
	}
}
