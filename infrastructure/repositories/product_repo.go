package repositories

import (
	"APIHEX_LPySP/config"
	"APIHEX_LPySP/domain"
)

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) GetProducts() ([]domain.Product, error) {
	var products []domain.Product
	result := config.DB.Find(&products)
	return products, result.Error
}

func (r *ProductRepository) CreateProduct(product *domain.Product) error {
	result := config.DB.Create(product)
	return result.Error
}

func (r *ProductRepository) GetProductByID(id uint) (*domain.Product, error) {
	var product domain.Product
	result := config.DB.First(&product, id)
	return &product, result.Error
}

func (r *ProductRepository) UpdateProduct(product *domain.Product) error {
	result := config.DB.Save(product)
	return result.Error
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	result := config.DB.Delete(&domain.Product{}, id)
	return result.Error
}
