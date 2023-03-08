package service

import (
	"sirawit/shop/internal/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *productService) GetProducts(id uint64) ([]model.Product, error) {
	products, err := p.db.GetProducts(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productService) CreateProduct(input model.Product) (*model.Product, error) {
	result, err := p.db.CreateProduct(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *productService) GetProduct(id uint64) (*model.Product, error) {
	product, err := p.db.GetProduct(id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "product not found")
	}
	return product, nil
}
