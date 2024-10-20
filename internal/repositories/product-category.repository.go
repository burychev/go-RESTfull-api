package repositories

import (
	"database/sql"
	"products/internal/models"
)

type IProductCategoryRepository interface {
	Create(data models.ProductCategory) (models.ProductCategory, error)
	Update(data models.ProductCategory) (models.ProductCategory, error)
	Delete(id int) error
	GetById(id int) (models.ProductCategory, error)
	GetAll() []models.ProductCategory
}

type ProductCategoryRepository struct {
	DB *sql.DB
}

func NewProductCategoryRepository(DB *sql.DB) IProductCategoryRepository {
	return &ProductCategoryRepository{DB: DB}
}

func (r *ProductCategoryRepository) Create(data models.ProductCategory) (models.ProductCategory, error) {
	var res models.ProductCategory
	err := r.DB.QueryRow("insert into ProductCategory (Name, Description) values ($1, $2) RETURNING Id",
		data.Name, data.Description).Scan(&res.Id)
	if err != nil {
		return models.ProductCategory{}, err
	}

	res.Name = data.Name
	res.Description = data.Description

	return res, nil
}

func (r *ProductCategoryRepository) GetAll() []models.ProductCategory {
	array := make([]models.ProductCategory, 0)
	rows, _ := r.DB.Query("SELECT Id, Name, Description FROM ProductCategory")

	defer rows.Close()

	for rows.Next() {
		var ProductCategory models.ProductCategory

		if err := rows.Scan(&ProductCategory.Id, &ProductCategory.Name, &ProductCategory.Description); err != nil {
			return nil
		}
		array = append(array, ProductCategory)
	}
	return array
}

func (r *ProductCategoryRepository) Update(data models.ProductCategory) (models.ProductCategory, error) {
	var res models.ProductCategory
	err := r.DB.QueryRow(`UPDATE ProductCategory SET Name = $2, Description = $3 WHERE Id = $1 RETURNING Id, Name, Description;`, data.Id, data.Name, data.Description).Scan(&res.Id, &res.Name, &res.Description)
	if err != nil {
		return models.ProductCategory{}, err
	}
	return res, nil
}

func (r *ProductCategoryRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM ProductCategory WHERE Id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductCategoryRepository) GetById(id int) (models.ProductCategory, error) {
	var res models.ProductCategory
	err := r.DB.QueryRow("SELECT Id, Name, Description from ProductCategory WHERE Id = $1",
		id).Scan(&res.Id, &res.Name, &res.Description)
	if err != nil {
		return models.ProductCategory{}, err
	}
	return res, nil
}
