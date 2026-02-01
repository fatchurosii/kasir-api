package service

import (
	"kasir-api/entity"
	"kasir-api/repository"
)

type ProductService struct {
	ProductRepo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepo: repo,
	}
}

func (s *ProductService) GetProductByID(id int) (*entity.Product, error) {
	return s.ProductRepo.GetProductByID(id)
}

func (s *ProductService) GetAllProductsService() ([]entity.Product, error) {
	return s.ProductRepo.GetAllProducts()
}

func (s *ProductService) CreateProduct(product *entity.Product) error {
	return s.ProductRepo.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(product *entity.Product) error {
	return s.ProductRepo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.ProductRepo.DeleteProduct(id)
}
