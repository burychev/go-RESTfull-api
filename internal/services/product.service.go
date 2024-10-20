package services

import (
	"products/internal/models"
	"products/internal/repositories"
)

type IProductService interface {
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(id int) error
	GetProductById(id int) (models.Product, error)
	GetAllProducts() []models.Product
}

type ProductService struct {
	repository repositories.IProductRepository
}

func NewProductService(r repositories.IProductRepository) IProductService {
	return &ProductService{repository: r}
}

func (s *ProductService) CreateProduct(product models.Product) (models.Product, error) {
	return s.repository.Create(product)
}

func (s *ProductService) UpdateProduct(product models.Product) (models.Product, error) {
	return s.repository.Update(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repository.Delete(id)
}

func (s *ProductService) GetProductById(id int) (models.Product, error) {
	return s.repository.GetById(id)
}

func (s *ProductService) GetAllProducts() []models.Product {
	return s.repository.GetAll()
}
