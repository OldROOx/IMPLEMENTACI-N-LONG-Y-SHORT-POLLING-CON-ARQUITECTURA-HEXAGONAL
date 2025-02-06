package application

import (
	"APIHEX_LPySP/domain"
	"APIHEX_LPySP/infrastructure/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) GetProducts() ([]domain.Product, error) {
	return s.productRepo.GetProducts()
}

func (s *ProductService) CreateProduct(product *domain.Product) error {
	return s.productRepo.CreateProduct(product)
}

func (s *ProductService) GetProductByID(id uint) (*domain.Product, error) {
	return s.productRepo.GetProductByID(id)
}

func (s *ProductService) UpdateProduct(product *domain.Product) error {
	return s.productRepo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.productRepo.DeleteProduct(id)
}
