package services

import (
	"products/internal/models"
	"products/internal/repositories"
)

type IProductCategoryService interface {
	CreateCategory(category models.ProductCategory) (models.ProductCategory, error)
	UpdateCategory(category models.ProductCategory) (models.ProductCategory, error)
	DeleteCategory(id int) error
	GetCategoryById(id int) (models.ProductCategory, error)
	GetAllCategories() []models.ProductCategory
}

type ProductCategoryService struct {
	repository repositories.IProductCategoryRepository
}

func NewProductCategoryService(r repositories.IProductCategoryRepository) IProductCategoryService {
	return &ProductCategoryService{repository: r}
}

func (s *ProductCategoryService) CreateCategory(category models.ProductCategory) (models.ProductCategory, error) {
	return s.repository.Create(category)
}

func (s *ProductCategoryService) UpdateCategory(category models.ProductCategory) (models.ProductCategory, error) {
	return s.repository.Update(category)
}

func (s *ProductCategoryService) DeleteCategory(id int) error {
	return s.repository.Delete(id)
}

func (s *ProductCategoryService) GetCategoryById(id int) (models.ProductCategory, error) {
	return s.repository.GetById(id)
}

func (s *ProductCategoryService) GetAllCategories() []models.ProductCategory {
	return s.repository.GetAll()
}
