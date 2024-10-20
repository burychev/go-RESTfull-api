package contracts

import "products/internal/models"

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	CategoryId  int     `json:"category_id" validate:"required"`
}

type CreateProductResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryId  int     `json:"category_id"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	CategoryId  int     `json:"category_id" validate:"required"`
}

type GetAllProductsResponse = []models.Product
