package repositories

import (
	"database/sql"
	"products/internal/models"
)

type IProductRepository interface {
	Create(data models.Product) (models.Product, error)
	Update(data models.Product) (models.Product, error)
	Delete(id int) error
	GetById(id int) (models.Product, error)
	GetAll() []models.Product
}

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(DB *sql.DB) IProductRepository {
	return &ProductRepository{DB: DB}
}

func (r *ProductRepository) Create(data models.Product) (models.Product, error) {
	var res models.Product
	err := r.DB.QueryRow("insert into Product (Name, Description, Price, CategoryId) values ($1, $2, $3, $4) RETURNING Id",
		data.Name, data.Description, data.Price, data.CategoryId).Scan(&res.Id)
	if err != nil {
		return models.Product{}, err
	}

	res.Name = data.Name
	res.Description = data.Description
	res.Price = data.Price
	res.CategoryId = data.CategoryId

	return res, nil
}

func (r *ProductRepository) Update(data models.Product) (models.Product, error) {
	var res models.Product
	err := r.DB.QueryRow(`UPDATE Product SET Name = $2, Description = $3, Price = $4, CategoryId = $5 WHERE Id = $1 RETURNING Id, Name, Description, Price, CategoryId;`, data.Id, data.Name, data.Description, data.Price, data.CategoryId).Scan(&res.Id, &res.Name, &res.Description, &res.Price, &res.CategoryId)
	if err != nil {
		return models.Product{}, err
	}
	return res, nil
}

func (r *ProductRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM Product WHERE Id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) GetById(id int) (models.Product, error) {
	var res models.Product
	err := r.DB.QueryRow("SELECT Id, Name, Description, Price, CategoryId from Product WHERE Id = $1",
		id).Scan(&res.Id, &res.Name, &res.Description, &res.Price, &res.CategoryId)
	if err != nil {
		return models.Product{}, err
	}
	return res, nil
}

func (r *ProductRepository) GetAll() []models.Product {
	array := make([]models.Product, 0)
	rows, _ := r.DB.Query("SELECT Id, Name, Description, Price, CategoryId FROM Product")

	defer rows.Close()

	for rows.Next() {
		var Product models.Product

		if err := rows.Scan(&Product.Id, &Product.Name, &Product.Description, &Product.Price, &Product.CategoryId); err != nil {
			return nil
		}
		array = append(array, Product)
	}
	return array
}
