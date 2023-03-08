package repository

import "sirawit/shop/internal/model"

const (
	limit = 2
)

func (p *productQuery) GetProducts(id uint64) ([]model.Product, error) {
	var products []model.Product
	if err := p.db.Where("id > ?", id).Limit(limit).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productQuery) GetProduct(id uint64) (*model.Product, error) {
	var product model.Product
	if err := p.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil

}

func (p *productQuery) CreateProduct(input model.Product) (*model.Product, error) {
	if err := p.db.Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}
