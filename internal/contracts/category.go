package contracts

import "products/internal/models"

type CreateProductCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateProductCategoryResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateProductCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type GetAllProductCategoryResponse = []models.ProductCategory
